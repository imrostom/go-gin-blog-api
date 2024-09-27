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
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	if !helpers.VerifyHashPassword(user.Password, c.PostForm("password")) {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, "pasword not macth")
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(user)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, "Could not generate token")
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": user, "token": token}, "Success")
}

func RegisterHandler(c *gin.Context) {
	// Validate form data
	messages, isValid := validations.AuthRegisterFormValidate(c)

	if !isValid {
		helpers.ErrorResponse(c, gin.H{"errors": messages}, "Error")
		return
	}

	// Create User via service
	user, err := services.CreateUser(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": user}, "Success")
}
