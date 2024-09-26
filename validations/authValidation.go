package validations

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/config"
	"github.com/imrostom/go-blog-api/models"
)

// Function to validate the entire form
func AuthLoginFormValidate(c *gin.Context) (map[string]string, bool) {
	messages := make(map[string]string)

	// Validate email
	if err := ValidateEmail(c.PostForm("email")); err != nil {
		messages["email"] = err.Error()
	}

	// Validate password
	if err := ValidatePassword(c.PostForm("password")); err != nil {
		messages["password"] = err.Error()
	}

	// Check if the form is valid
	isValid := len(messages) == 0

	return messages, isValid
}

// Function to validate the entire form
func AuthRegisterFormValidate(c *gin.Context) (map[string]string, bool) {
	messages := make(map[string]string)

	// Validate name
	if err := ValidateName(c.PostForm("name")); err != nil {
		messages["name"] = err.Error()
	}

	flagUniqueEmail := true

	// Validate email
	if err := ValidateEmail(c.PostForm("email")); err != nil {
		messages["email"] = err.Error()

		flagUniqueEmail = false
	}

	var existinguser models.User
	if flagUniqueEmail {
		if err := config.DB.Where("email = ?", c.PostForm("email")).First(&existinguser).Error; err == nil {
			messages["email"] = "user email must be unique"
		}
	}

	// Validate password
	if err := ValidatePassword(c.PostForm("password")); err != nil {
		messages["password"] = err.Error()
	}

	// Validate status
	status, _ := strconv.ParseUint(c.PostForm("status"), 10, 8)
	if err := ValidateStatus(uint8(status)); err != nil {
		messages["status"] = err.Error()
	}

	// Validate password
	if err := ValidateRole(c.PostForm("role")); err != nil {
		messages["role"] = err.Error()
	}

	// Check if the form is valid
	isValid := len(messages) == 0

	return messages, isValid
}
