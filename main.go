package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/config"
	"github.com/imrostom/go-blog-api/routes"
)

func init() {
	config.LoadEnvData()

	config.ConnectToDB()
}

func main() {
	// GIN initialize
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")

	// Setup routes
	routes.SetupRoutes(router)

	// Server listen and serve on 0.0.0.0:8080
	router.Run()
}
