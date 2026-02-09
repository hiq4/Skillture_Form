package server

// import (
// 	"log"
// 	"os"

// 	"github.com/gin-gonic/gin"
// )

// // Start initializes and starts the HTTP server
// func Start() {
// 	// Create a default Gin router
// 	r := gin.Default()

// 	// Setup middleware
// 	setupMiddleware(r)

// 	// Setup routes
// 	setupRoutes(r)

// 	// Get port from environment or default to 8080
// 	port := os.Getenv("SERVER_PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	log.Printf("Server running on port %s", port)
// 	r.Run(":" + port) // Run the server
// }
