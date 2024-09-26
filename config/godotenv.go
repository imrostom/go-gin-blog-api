package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvData() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
