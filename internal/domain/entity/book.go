package entity

import (
	"ddd-hexagonal-arch-go/internal/domain/valueobject"
	"fmt"
)

// Book is a struct that represents a book in the domain model.
type Book struct {
	ID       int
	Title    string
	Author   string
	Price    *valueobject.Price
	Quantity int
}

func NewBook(
	id int,
	title string,
	author string,
	price *valueobject.Price,
	quantity int) *Book {
	return &Book{
		ID:       id,
		Title:    title,
		Author:   author,
		Price:    price,
		Quantity: quantity,
	}
}

var ErrNotEnoughBooksStoked = fmt.Errorf("not enough stock for this book")

func (b *Book) ChechAvailability(quantity int) error {
	if b.Quantity < quantity {
		return ErrNotEnoughBooksStoked
	}

	return nil
}
