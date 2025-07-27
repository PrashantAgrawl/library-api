package books

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo BookRepository
}

func (h *Handler) GetAllBooks(c *gin.Context) {
	log.Println("Fetching all books")
	books, err := h.Repo.GetAllBooks()
	if err != nil {
		log.Printf("Error fetching books: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *Handler) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("Fetching book with ID %d", id)
	book, err := h.Repo.GetBookByID(id)
	if err != nil {
		log.Printf("Book not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) CreateBook(c *gin.Context) {
	log.Println("Creating a new book")
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		log.Printf("Invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.CreateBook(&book); err != nil {
		log.Printf("Error creating book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (h *Handler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("Updating book with ID %d", id)
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		log.Printf("Invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = id
	rowsAffected, err := h.Repo.UpdateBook(book)
	if err != nil {
		log.Printf("Error updating book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		log.Printf("Book with ID %d not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Printf("Attempting to delete book with ID %d", id)
	rowsAffected, err := h.Repo.DeleteBook(id)
	if err != nil {
		log.Printf("Error deleting book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		log.Printf("Book with ID %d not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
