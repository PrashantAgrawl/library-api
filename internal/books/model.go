package books

type Book struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    PublishedAt string `json:"published_at"`
}
