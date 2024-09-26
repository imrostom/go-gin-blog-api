package controllers

import (
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

func TestHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "emails/index.html", gin.H{
		"title": "Main website",
	})
}
