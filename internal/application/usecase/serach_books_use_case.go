package usecase

import (
	"ddd-hexagonal-arch-go/internal/domain/entity"
	"ddd-hexagonal-arch-go/internal/domain/ports/input"
	"ddd-hexagonal-arch-go/internal/domain/ports/output"
)

// SearchBooksUseCase is a struct that represents the use case for searching books.
type SearchBooksUseCase struct {
	BookRepository output.BookRepository
}

func NewSearchBooksUseCase(bookRepository output.BookRepository) input.SearchBooksInputPort {
	return &SearchBooksUseCase{
		BookRepository: bookRepository,
	}
}

// Search searches for books by title.
func (uc *SearchBooksUseCase) Search(query string) ([]*entity.Book, error) {
	return uc.BookRepository.FindByTitle(query)
}
