package routes

import (
	"github.com/dohyeoplim/blog-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	group := r.Group("/api/auth")
	{
		group.POST("/setup", controllers.SetupTOTP)
		group.POST("/verify", controllers.VerifyTOTP)
	}
}
