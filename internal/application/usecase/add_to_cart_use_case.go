package usecase

import (
	"ddd-hexagonal-arch-go/internal/domain/entity"
	"ddd-hexagonal-arch-go/internal/domain/ports/input"
	"ddd-hexagonal-arch-go/internal/domain/ports/output"
	"errors"
)

// AddCartUseCase is a struct that represents the shopping cart service. It has an AddToCart method that adds a book to the user's shopping cart.
type AddCartUseCase struct {
	BookRepository output.BookRepository
	CartRepository output.CartRepository
}

func NewAddCartUseCase(bookRepository output.BookRepository, cartRepository output.CartRepository) input.AddToCartInputPort {
	return &AddCartUseCase{
		BookRepository: bookRepository,
		CartRepository: cartRepository,
	}
}

// AddToCart adds a book to the user's shopping cart.
func (cs *AddCartUseCase) AddToCart(userID int, bookID int, quantity int) error {
	book, err := cs.BookRepository.FindByID(bookID)
	if err != nil {
		return err
	}

	err = book.ChechAvailability(quantity)
	if err != nil {
		return err
	}

	cart, err := cs.CartRepository.FindByUserID(userID)
	if err != nil && !errors.Is(err, output.ErrNotFound) {
		return err
	}
	
	if cart == nil {
		cart = entity.NewCart(userID, []*entity.CartItem{})
	}

	cart.AddItem(entity.NewCartItem(bookID, quantity))

	err = cs.CartRepository.Save(cart)
	if err != nil {
		return err
	}

	return nil
}
