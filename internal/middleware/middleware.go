package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger logs all incoming requests and response statuses
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()

		log.Printf("[%s] %s %s | %d | %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			statusCode,
			latency,
		)
	}
}

// TokenAuthMiddleware checks for a valid token in Authorization header
func TokenAuthMiddleware(validToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != validToken {
			log.Printf("Unauthorized request, token provided: %s", token)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized - invalid token",
			})
			return
		}
		c.Next()
	}
}
