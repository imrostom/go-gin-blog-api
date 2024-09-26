package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data gin.H, message string) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  true,
		"message": message,
	})
}

func ErrorResponse(c *gin.Context, data gin.H, message string) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  false,
		"message": message,
	})
}

func ValidationResponse(c *gin.Context, data gin.H, message string) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  false,
		"message": message,
	})
}
