package validations

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/config"
	"github.com/imrostom/go-blog-api/models"
)

// Struct to hold form data
type CategoryForm struct {
	Name   string
	Status uint8
}

// Function to validate the entire form
func CategoryFormValidate(c *gin.Context, isUpdate bool) (map[string]string, bool) {
	messages := make(map[string]string)

	// Validate name
	flagUniqueName := true

	if err := ValidateName(c.PostForm("name")); err != nil {
		messages["name"] = err.Error()

		flagUniqueName = false
	}

	var existingCategory models.Category
	if flagUniqueName {
		if isUpdate {
			id, _ := strconv.Atoi(c.Param("id"))

			if err := config.DB.Where("name = ? AND id != ?", c.PostForm("name"), id).First(&existingCategory).Error; err == nil {
				messages["name"] = "the category name must be unique"
			}
		} else {
			if err := config.DB.Where("name = ?", c.PostForm("name")).First(&existingCategory).Error; err == nil {
				messages["name"] = "the category name must be unique"
			}
		}

	}

	// Validate status
	status, _ := strconv.ParseUint(c.PostForm("status"), 10, 8)
	if err := ValidateStatus(uint8(status)); err != nil {
		messages["status"] = err.Error()
	}

	// Check if the form is valid
	isValid := len(messages) == 0

	return messages, isValid
}
