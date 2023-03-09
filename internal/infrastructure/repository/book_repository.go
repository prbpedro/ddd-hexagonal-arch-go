package repository

import (
	"ddd-hexagonal-arch-go/internal/domain/entity"
	"ddd-hexagonal-arch-go/internal/domain/ports/output"

	"errors"
)

// InMemoryBookRepository is a struct that represents an in-memory implementation of the BookRepository interface.
type InMemoryBookRepository struct {
	Books []*entity.Book
}

func NewInMemoryBookRepository(books []*entity.Book) output.BookRepository {
	return &InMemoryBookRepository{
		Books: books,
	}
}

// FindByID finds a book by ID.
func (r *InMemoryBookRepository) FindByID(id int) (*entity.Book, error) {
	for _, b := range r.Books {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, errors.New("book not found")
}

// FindByTitle finds books by title.
func (r *InMemoryBookRepository) FindByTitle(title string) ([]*entity.Book, error) {
	var results []*entity.Book
	for _, b := range r.Books {
		if b.Title == title {
			results = append(results, b)
		}
	}
	return results, nil
}

// Save saves a book.
func (r *InMemoryBookRepository) Save(book *entity.Book) error {
	if book.ID == 0 {
		book.ID = len(r.Books) + 1
		r.Books = append(r.Books, book)
	} else {
		for i, b := range r.Books {
			if b.ID == book.ID {
				r.Books[i] = book
				break
			}
		}
	}
	return nil
}
