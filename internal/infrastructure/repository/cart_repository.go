package repository

import (
	"ddd-hexagonal-arch-go/internal/domain/entity"
	"ddd-hexagonal-arch-go/internal/domain/ports/output"
)

// InMemoryCartRepository is a struct that represents an in-memory implementation of the CartRepository interface.
type InMemoryCartRepository struct {
	carts map[int]*entity.Cart
}

// NewInMemoryCartRepository creates a new instance of InMemoryCartRepository.
func NewInMemoryCartRepository() output.CartRepository {
	return &InMemoryCartRepository{
		carts: make(map[int]*entity.Cart),
	}
}

// FindByUserID finds a cart by user ID.
func (r *InMemoryCartRepository) FindByUserID(userID int) (*entity.Cart, error) {
	for _, cart := range r.carts {
		if cart.UserID == userID {
			return cart, nil
		}
	}
	return nil, output.ErrNotFound
}

// Save saves a domain.
func (r *InMemoryCartRepository) Save(cart *entity.Cart) error {
	r.carts[cart.UserID] = cart
	return nil
}
