package server

import "github.com/gin-gonic/gin"

// setupMiddleware configures all global middlewares
func setupMiddleware(r *gin.Engine) {
	// Example: Simple CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	})

	// TODO: Add JWT authentication middleware here
}
