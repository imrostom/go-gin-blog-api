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

// Fetch users with optional status filter
func GetUsers(c *gin.Context) ([]models.User, error) {
	var users []models.User

	status := c.Query("status")

	query := config.DB

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Scopes(helpers.Paginate(c)).Find(&users).Error; err != nil {
		return nil, errors.New("users not found")
	}

	return users, nil
}

// Create a new user
func CreateUser(c *gin.Context) (*models.User, error) {
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
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// Update user
func UpdateUser(c *gin.Context) (*models.User, error) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

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

	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Image = filePath
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	user.DateOfBirth = dateOfBirth
	user.Address = c.PostForm("adress")
	user.Password = hashedPassword
	user.Status = uint8(status)
	user.Role = c.PostForm("role")

	if err := config.DB.Save(user).Error; err != nil {
		return nil, errors.New("failed to update user")
	}

	return &user, nil
}

// Delete user by ID
func DeleteUser(id uint) error {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
