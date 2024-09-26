package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSettingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Get Setting")
}

func CreateSettingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Create Setting")
}

func ShowSettingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Show Setting")
}

func UpdateSettingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Update Setting")
}

func DeleteSettingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Delete Setting")
}
