package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/config"
	"github.com/imrostom/go-blog-api/helpers"
	"github.com/imrostom/go-blog-api/models"
)

func RegisterUser(c *gin.Context) (*models.User, error) {
	status, _ := strconv.ParseUint(c.PostForm("status"), 10, 8)
	dateOfBirth, _ := time.Parse("02-01-2006", c.PostForm("date_of_birth"))
	hashedPassword, _ := helpers.HashPassword(c.PostForm("password"))

	// Image upload logic (optional)
	file, err := c.FormFile("image")
	var filePath string // This will store the image path

	if err == nil {
		// If an image is uploaded, handle the image upload using the service
		filePath, err = helpers.UploadImage(c, file, "./uploads")
		if err != nil {
			return nil, err // Return the error from the service
		}
	} else {
		// No image uploaded, set default value (optional: use a placeholder image or keep it empty)
		filePath = ""
	}

	user := &models.User{
		Name:        c.PostForm("name"),
		Image:       filePath,
		Email:       c.PostForm("email"),
		Phone:       c.PostForm("phone"),
		DateOfBirth: dateOfBirth,
		Address:     c.PostForm("adress"),
		Password:    hashedPassword,
		Status:      uint8(status),
		Role:        c.PostForm("role"),
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return user, nil
}

// Fetch user by ID
func AuthUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// Fetch user by ID
func AuthUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
