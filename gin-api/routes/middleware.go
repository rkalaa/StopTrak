package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs each incoming request
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log details after request is processed
		duration := time.Since(start)
		log.Printf("[%s] %s %d %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}

// AuthMiddleware validates an authentication token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !validateToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}

// validateToken is a placeholder for actual token validation logic
func validateToken(token string) bool {
	// Replace with actual token validation logic
	return token == "valid-token"
}
