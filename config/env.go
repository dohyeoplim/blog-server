package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("RAILWAY_ENVIRONMENT_NAME") == "production" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("No .env file found. Using system environment variables.")
		} else {
			fmt.Println(".env file loaded successfully.")
		}
	}
}
