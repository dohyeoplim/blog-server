package routes

import (
	"github.com/dohyeoplim/blog-server/controllers"
	"github.com/dohyeoplim/blog-server/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {
	group := r.Group("/api/posts")
	{
		group.GET("", controllers.GetAllPosts)
		group.GET("/:id", controllers.GetPost)

		group.Use(middleware.JWTAuthMiddleware())
		{
			group.POST("", controllers.CreatePost)
			group.PUT("/:id", controllers.UpdatePost)
			group.DELETE("/:id", controllers.DeletePost)
		}
	}
}
