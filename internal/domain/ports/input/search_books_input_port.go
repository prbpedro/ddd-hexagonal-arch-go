package input

import "ddd-hexagonal-arch-go/internal/domain/entity"

// SearchBooksInputPort is an interface that defines the input port for searching books.
type SearchBooksInputPort interface {
	Search(query string) ([]*entity.Book, error)
}
