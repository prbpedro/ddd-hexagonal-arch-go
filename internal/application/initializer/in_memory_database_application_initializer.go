package initializer

import (
	"ddd-hexagonal-arch-go/internal/application/usecase"
	"ddd-hexagonal-arch-go/internal/domain/entity"
	"ddd-hexagonal-arch-go/internal/domain/ports/initializer"
	"ddd-hexagonal-arch-go/internal/domain/ports/input"
	"ddd-hexagonal-arch-go/internal/domain/valueobject"
	"ddd-hexagonal-arch-go/internal/infrastructure/repository"
)

type InMemoryDatabaseApplicationInitializer struct {
}

func NewInMemoryDatabaseApplicationInitializer() initializer.ApplicationInitializer {
	return &InMemoryDatabaseApplicationInitializer{}
}

func (s *InMemoryDatabaseApplicationInitializer) InitApp() (input.SearchBooksInputPort, input.AddToCartInputPort, error) {

	books := []*entity.Book{
		entity.NewBook(1, "Book 1", "Author 1", valueobject.NewPrice(10), 5),
		entity.NewBook(1, "Book 1", "Author 1", valueobject.NewPrice(15), 3),
		entity.NewBook(1, "Book 1", "Author 1", valueobject.NewPrice(20), 2),
	}

	// Initialize repositories
	bookRepository := repository.NewInMemoryBookRepository(books)
	cartRepository := repository.NewInMemoryCartRepository()

	// Initialize use cases
	searchBooksUseCase := usecase.NewSearchBooksUseCase(bookRepository)

	// Initialize services
	addCartUseCase := usecase.NewAddCartUseCase(bookRepository, cartRepository)

	return searchBooksUseCase, addCartUseCase, nil
}
