package main

import (
	// Adjust based on your module name
	// If you need to import models
	// If you have a services directory
	// If you have a utils directory
	"gin-dep/config"
	"gin-dep/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to the database
	err = db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Set up API routes
	routes.SetupRoutes(router)

	// Start server
	log.Println("Server starting on port:", config.GetPort())
	router.Run(":" + config.GetPort())
}
