package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/popova-mariia/cvwo-web-forum/internal/controllers"
	"github.com/popova-mariia/cvwo-web-forum/internal/db"
)

func main() {
	//db.InitDB("user=username dbname=yourdbname sslmode=disable")
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	// Optionally, run migrations to ensure tables are created
	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v\n", err)
	}

	r := gin.Default()

	// Enable CORS for all origins
	r.Use(cors.Default())

	// Define your routes
	r.POST("/api/threads", controllers.CreateThread)

	// Start the server on port 8080
	r.Run(":8080")
}
