package main

import (
	"log"
	"os"

	"github.com/dohyeoplim/blog-server/config"
	"github.com/dohyeoplim/blog-server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()
	routes.RegisterAuthRoutes(r)
	routes.RegisterPostRoutes(r)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running at http://localhost:%s", port)
	r.Run(":" + port)
}
