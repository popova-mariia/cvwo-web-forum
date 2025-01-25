package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/popova-mariia/cvwo-web-forum/internal/db"
	"github.com/popova-mariia/cvwo-web-forum/internal/router"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}
	r := gin.Default()

	// Custom CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.SetupRouter(r)

	r.Run(":8080")
}
