package controllers

import (
	"log"
	"net/http"
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
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
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
	user, err := services.CreateUser(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	// Load data to pass into the template
	data := UserData{
		AppName:          "EasyBlog", // Replace with actual data
		VerificationLink: "https://wwww.google.com",
		UserName:         user.Name,
	}

	// Parse the HTML email template
	body, err := helpers.RenderTemplate("templates/mails/user-registration.html", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to load email template"})
		return
	}

	// Send the email
	err = helpers.SendEmail(user.Email, "Welcome to the service", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send email"})
		log.Println("Error sending email:", err)
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": user}, "Success")
}

func ShowUserHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": user}, "Success")
}

func UpdateUserHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := services.GetUserByID(uint(id))
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
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
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": user}, "Success")
}

func DeleteUserHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Delete User via service
	if err := services.DeleteUser(uint(id)); err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"user": gin.H{}}, "Success")
}
