package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/helpers"
	"github.com/imrostom/go-blog-api/services"
	"github.com/imrostom/go-blog-api/validations"
)

func GetCategoryHandler(c *gin.Context) {
	// Fetch categories from the service
	categories, err := services.GetCategories(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	helpers.SuccessResponse(c, gin.H{"categories": categories, "page": page}, "Success")
}

func CreateCategoryHandler(c *gin.Context) {
	// Validate form data
	messages, isValid := validations.CategoryFormValidate(c, false)
	fmt.Println("ddd")
	if !isValid {
		helpers.ErrorResponse(c, gin.H{"errors": messages}, "Error")
		return
	}

	// Create category via service
	category, err := services.CreateCategory(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"category": category}, "Success")
}

func ShowCategoryHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := services.GetCategoryByID(uint(id))
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"category": category}, "Success")
}

func UpdateCategoryHandler(c *gin.Context) {
	// Validate form data
	messages, isValid := validations.CategoryFormValidate(c, true)
	if !isValid {
		helpers.ErrorResponse(c, gin.H{"errors": messages}, "Error")
		return
	}

	// Update category via service
	category, err := services.UpdateCategory(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"category": category}, "Success")
}

func DeleteCategoryHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Delete category via service
	if err := services.DeleteCategory(uint(id)); err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"category": gin.H{}}, "Success")
}
