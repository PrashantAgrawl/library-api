package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTokenAuthMiddleware_ValidToken(t *testing.T) {
	router := gin.New()
	router.Use(TokenAuthMiddleware("test-token"))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	req, _ := http.NewRequest("GET", "/ping", nil)
	req.Header.Set("Authorization", "test-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestTokenAuthMiddleware_InvalidToken(t *testing.T) {
	router := gin.New()
	router.Use(TokenAuthMiddleware("test-token"))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	req, _ := http.NewRequest("GET", "/ping", nil)
	req.Header.Set("Authorization", "wrong-token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestTokenAuthMiddleware_NoToken(t *testing.T) {
	router := gin.New()
	router.Use(TokenAuthMiddleware("test-token"))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	req, _ := http.NewRequest("GET", "/ping", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}
