package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/config"
	"github.com/imrostom/go-blog-api/helpers"
	"github.com/imrostom/go-blog-api/models"
)

// Fetch posts with optional status filter
func GetPosts(c *gin.Context) ([]models.Post, error) {
	var posts []models.Post

	status := c.Query("status")

	query := config.DB

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Scopes(helpers.Paginate(c)).Find(&posts).Error; err != nil {
		return nil, errors.New("posts not found")
	}

	return posts, nil
}

// Create a new post
func CreatePost(c *gin.Context) (*models.Post, error) {
	// Safely retrieve and assert userId from c.Keys
	userIdValue := c.Keys["userId"]
	userIdStr := userIdValue.(string)
	userId, _ := strconv.Atoi(userIdStr)

	categoryId, _ := strconv.Atoi(c.PostForm("category_id"))
	status, _ := strconv.ParseUint(c.PostForm("status"), 10, 8)
	PublishedAt, _ := time.Parse("02-01-2006", c.PostForm("published_at"))

	// Image upload logic (optional)
	file, err := c.FormFile("image")
	var filePath string // This will store the image path

	if err == nil {
		// If an image is uploaded, handle the image upload using the service
		filePath, err = helpers.UploadImage(c, file, "./uploads")
		if err != nil {
			return nil, err // Return the error from the service
		}
	} else {
		// No image uploaded, set default value (optional: use a placeholder image or keep it empty)
		filePath = ""
	}

	post := &models.Post{
		Title:       c.PostForm("title"),
		Slug:        helpers.GenerateSlug(c.PostForm("title")),
		UserId:      uint(userId),
		CategoryId:  uint(categoryId),
		Content:     c.PostForm("content"),
		Image:       filePath,
		Status:      uint8(status),
		PublishedAt: PublishedAt,
	}

	if err := config.DB.Create(post).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return post, nil
}

// Fetch post by ID
func GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return nil, errors.New("post not found")
	}

	return &post, nil
}

// Update post
func UpdatePost(c *gin.Context) (*models.Post, error) {
	id, _ := strconv.Atoi(c.Param("id"))

	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return nil, errors.New("post not found")
	}

	// Safely retrieve and assert userId from c.Keys
	userIdValue := c.Keys["userId"]
	userIdStr := userIdValue.(string)
	userId, _ := strconv.Atoi(userIdStr)

	categoryId, _ := strconv.Atoi(c.PostForm("category_id"))
	status, _ := strconv.ParseUint(c.PostForm("status"), 10, 8)
	publishedAt, _ := time.Parse("02-01-2006", c.PostForm("published_at"))

	// Image upload logic (optional)
	file, err := c.FormFile("image")
	var filePath string // This will store the image path

	if err == nil {
		// If an image is uploaded, handle the image upload using the service
		filePath, err = helpers.UploadImage(c, file, "./uploads")
		if err != nil {
			return nil, err // Return the error from the service
		}
	} else {
		// No image uploaded, set default value (optional: use a placeholder image or keep it empty)
		filePath = ""
	}

	post.ID = uint(id)
	post.Title = c.PostForm("title")
	post.Slug = c.PostForm("title")
	post.UserId = uint(userId)
	post.CategoryId = uint(categoryId)
	post.Content = c.PostForm("content")
	post.Image = filePath
	post.Status = uint8(status)
	post.PublishedAt = publishedAt

	if err := config.DB.Save(post).Error; err != nil {
		return nil, errors.New("failed to update post")
	}

	return &post, nil
}

// Delete post by ID
func DeletePost(id uint) error {
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return errors.New("post not found")
	}

	if err := config.DB.Delete(&post, id).Error; err != nil {
		return errors.New("failed to delete post")
	}

	return nil
}
