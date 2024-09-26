package services

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/config"
	"github.com/imrostom/go-blog-api/helpers"
	"github.com/imrostom/go-blog-api/models"
)

// Fetch categories with optional status filter
func GetCategories(c *gin.Context) ([]models.Category, error) {
	status := c.Query("status")

	var categories []models.Category

	query := config.DB
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Scopes(helpers.Paginate(c)).Find(&categories).Error; err != nil {
		return nil, errors.New("categories not found")
	}

	return categories, nil
}

// Create a new category
func CreateCategory(c *gin.Context) (*models.Category, error) {
	status, _ := strconv.ParseUint(c.PostForm("status"), 10, 8)

	category := &models.Category{
		Name:   c.PostForm("name"),
		Status: uint8(status),
	}

	if err := config.DB.Create(category).Error; err != nil {
		return nil, errors.New("failed to create category")
	}

	return category, nil
}

// Fetch category by ID
func GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return nil, errors.New("category not found")
	}

	return &category, nil
}

// Update category
func UpdateCategory(c *gin.Context) (*models.Category, error) {
	id, _ := strconv.Atoi(c.Param("id"))

	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return nil, errors.New("category not found")
	}

	status, _ := strconv.ParseUint(c.PostForm("status"), 10, 8)

	category.ID = uint(id)
	category.Name = c.PostForm("name")
	category.Status = uint8(status)

	if err := config.DB.Save(&category).Error; err != nil {
		return nil, errors.New("failed to update category")
	}

	return &category, nil
}

// Delete category by ID
func DeleteCategory(id uint) error {
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return errors.New("category not found")
	}

	if err := config.DB.Delete(&category, id).Error; err != nil {
		return errors.New("failed to delete category")
	}

	return nil
}
