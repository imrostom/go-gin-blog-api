package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/helpers"
	"github.com/imrostom/go-blog-api/services"
	"github.com/imrostom/go-blog-api/validations"
)

func GetUserHandler(c *gin.Context) {
	// Fetch users from the service
	users, err := services.GetUsers(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"User": gin.H{}}, err.Error())
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	helpers.SuccessResponse(c, gin.H{"users": users, "page": page}, "Success")
}

func CreateUserHandler(c *gin.Context) {
	// Validate form data
	messages, isValid := validations.UserFormValidate(c, false)

	if !isValid {
		helpers.ErrorResponse(c, gin.H{"errors": messages}, "Error")
		return
	}

	// Create User via service
	User, err := services.CreateUser(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{"name": err.Error()}}, "Error")
		return
	}

	helpers.SuccessResponse(c, gin.H{"User": User}, "Success")
}

func ShowUserHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"user": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": user}, "Success")
}

func UpdateUserHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := services.GetUserByID(uint(id))
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"user": gin.H{}}, err.Error())
		return
	}

	// Validate form data
	messages, isValid := validations.UserFormValidate(c, true)
	if !isValid {
		helpers.ErrorResponse(c, gin.H{"errors": messages}, "Error")
		return
	}

	// Update User via service
	user, err := services.UpdateUser(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{"name": err.Error()}}, "Error")
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": user}, "Success")
}

func DeleteUserHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Delete User via service
	if err := services.DeleteUser(uint(id)); err != nil {
		helpers.ErrorResponse(c, gin.H{"user": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": gin.H{}}, "Success")
}
