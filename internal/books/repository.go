package books

import "database/sql"

type BookRepository interface {
	GetAllBooks() ([]Book, error)
	GetBookByID(id int) (Book, error)
	CreateBook(b *Book) error
	UpdateBook(b Book) (int64, error)
	DeleteBook(id int) (int64, error)
}

type SQLRepository struct {
	DB *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{DB: db}
}

func (r *SQLRepository) GetAllBooks() ([]Book, error) {
	rows, err := r.DB.Query("SELECT id, title, author, published_at FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.PublishedAt); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *SQLRepository) GetBookByID(id int) (Book, error) {
	var b Book
	err := r.DB.QueryRow("SELECT id, title, author, published_at FROM books WHERE id=$1", id).
		Scan(&b.ID, &b.Title, &b.Author, &b.PublishedAt)
	return b, err
}

func (r *SQLRepository) CreateBook(b *Book) error {
	return r.DB.QueryRow(
		"INSERT INTO books (title, author, published_at) VALUES ($1, $2, $3) RETURNING id",
		b.Title, b.Author, b.PublishedAt,
	).Scan(&b.ID)
}

func (r *SQLRepository) UpdateBook(b Book) (int64, error) {
	result, err := r.DB.Exec(
		"UPDATE books SET title=$1, author=$2, published_at=$3 WHERE id=$4",
		b.Title, b.Author, b.PublishedAt, b.ID,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, nil
}

func (r *SQLRepository) DeleteBook(id int) (int64, error) {
	result, err := r.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, nil
}
