package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/helpers"
)

func DefaultHandler(c *gin.Context) {
	// Prepare response data
	responseData := gin.H{
		"content": "This is home page",
	}

	// Send the JSON response
	helpers.SuccessResponse(c, responseData, "Success")
}

type UserData struct {
	AppName          string
	VerificationLink string
	UserName         string
}

func TestHandler(c *gin.Context) {
	// Load data to pass into the template
	data := UserData{
		AppName:          "EasyBlog", // Replace with actual data
		VerificationLink: "https://wwww.google.com",
		UserName:         "dsss",
	}

	// Parse the HTML email template
	body, err := helpers.RenderTemplate("templates/mails/user-registration.html", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to load email template"})
		return
	}

	// Send the email
	err = helpers.SendEmail("recipient@example.com", "Welcome to the service", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send email"})
		log.Println("Error sending email:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
