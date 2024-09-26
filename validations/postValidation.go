package validations

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/config"
	"github.com/imrostom/go-blog-api/models"
)

// Struct to hold form data
type PostForm struct {
	Name   string
	Status uint8
}

// Function to validate the entire form
func PostFormValidate(c *gin.Context, isUpdate bool) (map[string]string, bool) {
	messages := make(map[string]string)

	flagUniqueTitle := true
	// Validate title
	if err := ValidateTitle(c.PostForm("title")); err != nil {
		messages["title"] = err.Error()

		flagUniqueTitle = false
	}

	var existingPost models.Post
	if flagUniqueTitle {
		if isUpdate {
			id, _ := strconv.Atoi(c.Param("id"))

			if err := config.DB.Where("title = ? AND id != ?", c.PostForm("title"), id).First(&existingPost).Error; err == nil {
				messages["title"] = "Post title must be unique"
			}
		} else {
			if err := config.DB.Where("title = ?", c.PostForm("title")).First(&existingPost).Error; err == nil {
				messages["title"] = "Post title must be unique"
			}
		}

	}

	// Validate category
	if err := ValidateCategory(c.PostForm("category_id")); err != nil {
		messages["categrory"] = err.Error()
	}

	// Validate content
	if err := ValidateContent(c.PostForm("content")); err != nil {
		messages["content"] = err.Error()
	}

	// Validate status
	status, _ := strconv.ParseUint(c.PostForm("status"), 10, 8)
	if err := ValidateStatus(uint8(status)); err != nil {
		messages["status"] = err.Error()
	}

	// Validate published_at
	if err := ValidateDate(c.PostForm("published_at")); err != nil {
		messages["published_at"] = err.Error()
	}

	// Check if the form is valid
	isValid := len(messages) == 0

	return messages, isValid
}
