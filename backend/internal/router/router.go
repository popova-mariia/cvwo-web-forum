package router

import (
	"github.com/gin-gonic/gin"
	"github.com/popova-mariia/cvwo-web-forum/internal/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id/score", controllers.UpdateScore)
	r.PUT("/users/:id/anonymous", controllers.ToggleAnonymous)

	// Thread routes
	r.POST("/threads", controllers.CreateThread)
	r.GET("/threads", controllers.ListThreads)
	r.POST("/threads/:thread_id/posts", controllers.CreatePost)
	r.GET("/threads/:thread_id/posts", controllers.ListPosts)
	r.GET("/threads/:thread_id/summary", controllers.GetThreadSummary)

	// Example route for "share"
	r.GET("/threads/:thread_id/share", func(c *gin.Context) {
		// In reality, you might create a short link or do something more advanced
		threadID := c.Param("thread_id")
		shareURL := "https://myforum.example.com/threads/" + threadID
		c.JSON(200, gin.H{"share_link": shareURL})
	})

	// Example route for "notifications" - placeholder
	r.GET("/notifications/:user_id", func(c *gin.Context) {
		// You’d fetch from your notifications table or an external service
		// For now, we just return a dummy message
		c.JSON(200, gin.H{
			"notifications": []string{"You have 3 new replies", "New poll in Action Movies"},
		})
	})

	return r
}
