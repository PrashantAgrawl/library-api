package router

import (
	"database/sql"
	"library-api/internal/books"
	"library-api/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Request logger
	r.Use(middleware.RequestLogger())

	// health check endpoint (no token required)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Apply token authentication for all other routes
	protected := r.Group("/")
	protected.Use(middleware.TokenAuthMiddleware(getAuthToken()))

	repo := books.NewSQLRepository(db)
	h := books.Handler{Repo: repo}

	protected.GET("/books", h.GetAllBooks)
	protected.GET("/books/:id", h.GetBook)
	protected.POST("/books", h.CreateBook)
	protected.PUT("/books/:id", h.UpdateBook)
	protected.DELETE("/books/:id", h.DeleteBook)

	return r
}

func getAuthToken() string {
	token := os.Getenv("API_TOKEN")
	if token == "" {
		return "my-secret-token"
	}
	return token
}
