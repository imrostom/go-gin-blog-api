package main

import (
	"fmt"

	"github.com/imrostom/go-blog-api/config"
	"github.com/imrostom/go-blog-api/models"
)

func init() {
	config.LoadEnvData()

	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(
		&models.Setting{},
		&models.User{},
		&models.Category{},
		&models.Post{},
	)

	fmt.Println("Successfully run migration")
}
