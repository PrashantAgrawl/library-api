package books

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mockRepository implements BookRepository for testing
type mockRepository struct {
	books []Book
}

func (m *mockRepository) GetAllBooks() ([]Book, error) {
	return m.books, nil
}

func (m *mockRepository) GetBookByID(id int) (Book, error) {
	for _, b := range m.books {
		if b.ID == id {
			return b, nil
		}
	}
	return Book{}, ErrNotFound
}

func (m *mockRepository) CreateBook(b *Book) error {
	b.ID = len(m.books) + 1
	m.books = append(m.books, *b)
	return nil
}

func (m *mockRepository) UpdateBook(b Book) (int64, error) {
	for i, book := range m.books {
		if book.ID == b.ID {
			m.books[i] = b
			return 1, nil
		}
	}
	return 0, nil
}

func (m *mockRepository) DeleteBook(id int) (int64, error) {
	for i, book := range m.books {
		if book.ID == id {
			m.books = append(m.books[:i], m.books[i+1:]...)
			return 1, nil
		}
	}
	return 0, nil
}

// ErrNotFound simulates a "book not found" error
var ErrNotFound = gin.Error{Err: http.ErrNoLocation, Type: gin.ErrorTypePublic}

func setupTestRouter(repo BookRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	h := Handler{Repo: repo}
	r := gin.Default()

	r.GET("/books", h.GetAllBooks)
	r.GET("/books/:id", h.GetBook)
	r.POST("/books", h.CreateBook)
	r.PUT("/books/:id", h.UpdateBook)
	r.DELETE("/books/:id", h.DeleteBook)

	return r
}

func TestGetAllBooks(t *testing.T) {
	repo := &mockRepository{
		books: []Book{{ID: 1, Title: "Book A", Author: "Author A", PublishedAt: "2025"}},
	}
	r := setupTestRouter(repo)

	req, _ := http.NewRequest("GET", "/books", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var books []Book
	json.Unmarshal(resp.Body.Bytes(), &books)
	assert.Len(t, books, 1)
	assert.Equal(t, "Book A", books[0].Title)
}

func TestGetBook_Success(t *testing.T) {
	repo := &mockRepository{
		books: []Book{{ID: 1, Title: "Book A", Author: "Author A", PublishedAt: "2025"}},
	}
	r := setupTestRouter(repo)

	req, _ := http.NewRequest("GET", "/books/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var book Book
	json.Unmarshal(resp.Body.Bytes(), &book)
	assert.Equal(t, 1, book.ID)
}

func TestGetBook_NotFound(t *testing.T) {
	repo := &mockRepository{
		books: []Book{{ID: 1, Title: "Book A", Author: "Author A", PublishedAt: "2025"}},
	}
	r := setupTestRouter(repo)

	req, _ := http.NewRequest("GET", "/books/999", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestCreateBook(t *testing.T) {
	repo := &mockRepository{}
	r := setupTestRouter(repo)

	newBook := `{"title":"New Book","author":"New Author","published_at":"2025"}`
	req, _ := http.NewRequest("POST", "/books", bytes.NewBufferString(newBook))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var book Book
	json.Unmarshal(resp.Body.Bytes(), &book)
	assert.Equal(t, "New Book", book.Title)
	assert.Equal(t, 1, book.ID)
}

func TestUpdateBook_Success(t *testing.T) {
	repo := &mockRepository{
		books: []Book{{ID: 1, Title: "Book A", Author: "Author A", PublishedAt: "2025"}},
	}
	r := setupTestRouter(repo)

	updatedBook := `{"title":"Updated Book","author":"Updated Author","published_at":"2026"}`
	req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBufferString(updatedBook))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var book Book
	json.Unmarshal(resp.Body.Bytes(), &book)
	assert.Equal(t, "Updated Book", book.Title)
}

func TestUpdateBook_NotFound(t *testing.T) {
	repo := &mockRepository{
		books: []Book{{ID: 1, Title: "Book A", Author: "Author A", PublishedAt: "2025"}},
	}
	r := setupTestRouter(repo)

	updatedBook := `{"title":"Updated Book","author":"Updated Author","published_at":"2026"}`
	req, _ := http.NewRequest("PUT", "/books/999", bytes.NewBufferString(updatedBook))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestDeleteBook_Success(t *testing.T) {
	repo := &mockRepository{
		books: []Book{{ID: 1, Title: "Book A", Author: "Author A", PublishedAt: "2025"}},
	}
	r := setupTestRouter(repo)

	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// Ensure book list is now empty
	assert.Len(t, repo.books, 0)
}

func TestDeleteBook_NotFound(t *testing.T) {
	repo := &mockRepository{
		books: []Book{{ID: 1, Title: "Book A", Author: "Author A", PublishedAt: "2025"}},
	}
	r := setupTestRouter(repo)

	req, _ := http.NewRequest("DELETE", "/books/999", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
