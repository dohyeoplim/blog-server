package main

import (
	"log"
	"os"
	"time"

	"github.com/dohyeoplim/blog-server/config"
	"github.com/dohyeoplim/blog-server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://www.dohyeoplim.me", "https://dohyeoplim.me"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterAuthRoutes(r)
	routes.RegisterPostRoutes(r)
	routes.RegisterUploadRoutes(r)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running at http://localhost:%s", port)
	r.Run(":" + port)
}
