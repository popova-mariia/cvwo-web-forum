package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/popova-mariia/cvwo-web-forum/internal/db"
	"github.com/popova-mariia/cvwo-web-forum/internal/router"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	r := router.SetupRouter()
	r.Use(cors.Default())
	r.Run(":8080")
}
