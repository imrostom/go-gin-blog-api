package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/helpers"
	"github.com/imrostom/go-blog-api/services"
	"github.com/imrostom/go-blog-api/validations"
)

func GetPostHandler(c *gin.Context) {
	// Fetch posts from the service
	posts, err := services.GetPosts(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"post": gin.H{}}, err.Error())
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	helpers.SuccessResponse(c, gin.H{"posts": posts, "page": page}, "Success")
}

func CreatePostHandler(c *gin.Context) {
	// Validate form data
	messages, isValid := validations.PostFormValidate(c, false)

	if !isValid {
		helpers.ErrorResponse(c, gin.H{"errors": messages}, "Error")
		return
	}

	// Create Post via service
	post, err := services.CreatePost(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"post": post}, "Success")
}

func ShowPostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := services.GetPostByID(uint(id))
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"post": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"post": post}, "Success")
}

func UpdatePostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := services.GetPostByID(uint(id))
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"post": gin.H{}}, err.Error())
		return
	}

	// Validate form data
	messages, isValid := validations.PostFormValidate(c, true)
	if !isValid {
		helpers.ErrorResponse(c, gin.H{"errors": messages}, "Error")
		return
	}

	// Update post via service
	post, err := services.UpdatePost(c)
	if err != nil {
		helpers.ErrorResponse(c, gin.H{"errors": gin.H{"name": err.Error()}}, "Error")
		return
	}

	helpers.SuccessResponse(c, gin.H{"post": post}, "Success")
}

func DeletePostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Delete Post via service
	if err := services.DeletePost(uint(id)); err != nil {
		helpers.ErrorResponse(c, gin.H{"post": gin.H{}}, err.Error())
		return
	}

	helpers.SuccessResponse(c, gin.H{"post": gin.H{}}, "Success")
}
