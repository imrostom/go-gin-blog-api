package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/helpers"
	"github.com/imrostom/go-blog-api/services"
	"github.com/imrostom/go-blog-api/validations"
)

func LoginHandler(c *gin.Context) {
	// Validate form data
	messages, isValid := validations.AuthLoginFormValidate(c)

	if !isValid {
		helpers.ErrorResponse(c, gin.H{"errors": messages}, "Error")
		return
	}

	user, err := services.AuthUserByEmail(c.PostForm("email"))
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"user": gin.H{}}, err.Error())
		return
	}

	if !helpers.VerifyHashPassword(user.Password, c.PostForm("password")) {
		helpers.ErrorResponse(c, gin.H{"user": gin.H{}}, "pasword not macth")
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(user)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"user": gin.H{}}, "Could not generate token")
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": user, "token": token}, "Success")
}

func RegisterHandler(c *gin.Context) {
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
