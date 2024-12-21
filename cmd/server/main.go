package main

import (
	"fmt"
	"log"

	"github.com/popova-mariia/cvwo-web-forum/internal/db"
	"github.com/popova-mariia/cvwo-web-forum/internal/router"
)

func main() {
	// Initialize database
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.DB.Close()

	// Migrate / ensure tables
	err = db.Migrate()
	if err != nil {
		log.Fatalf("Failed to migrate DB: %v", err)
	}

	// Setup router
	r := router.SetupRouter()

	fmt.Println("Starting Go Movie Forum on http://localhost:8080")
	r.Run(":8080") // Listen and serve on port 8080
}
