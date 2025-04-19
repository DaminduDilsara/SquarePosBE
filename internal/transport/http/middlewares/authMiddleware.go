package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware - Middleware to check for auth token in the headers
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the "Authorization" header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			c.Abort() // Stop further processing
			return
		}

		// Continue processing the request
		c.Next()
	}
}
