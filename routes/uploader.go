package routes

import (
	"github.com/dohyeoplim/blog-server/controllers"
	"github.com/dohyeoplim/blog-server/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUploadRoutes(r *gin.Engine) {
	group := r.Group("/api/upload")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("", controllers.UploadImage)
	}

}
