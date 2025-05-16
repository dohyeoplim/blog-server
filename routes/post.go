package routes

import (
	"github.com/dohyeoplim/blog-server/controllers"
	"github.com/dohyeoplim/blog-server/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {
	group := r.Group("/api/posts")
	{
		group.GET("", middleware.OptionalJWTAuthMiddleware(), controllers.GetAllPosts)
		group.GET("/:slug", middleware.OptionalJWTAuthMiddleware(), controllers.GetPost)

		group.Use(middleware.JWTAuthMiddleware())
		{
			group.POST("", controllers.CreatePost)
			group.PATCH("/:slug", controllers.UpdatePost)
			group.DELETE("/:slug", controllers.DeletePost)
		}
	}
}
