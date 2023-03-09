package output

import "ddd-hexagonal-arch-go/internal/domain/entity"

// BookRepository is an interface that defines the output port for storing and retrieving books.
type BookRepository interface {
	FindByID(id int) (*entity.Book, error)
	FindByTitle(title string) ([]*entity.Book, error)
	Save(book *entity.Book) error
}
