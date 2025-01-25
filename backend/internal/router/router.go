package router

import (
	"github.com/gin-gonic/gin"
	"github.com/popova-mariia/cvwo-web-forum/internal/controllers"
	"github.com/popova-mariia/cvwo-web-forum/internal/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/users", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/threads", controllers.CreateThread)
		protected.GET("/threads", controllers.ListThreads)
		protected.POST("/threads/:thread_id/posts", controllers.CreatePost)
		protected.GET("/threads/:thread_id/posts", controllers.ListPosts)
	}

	return r
}
