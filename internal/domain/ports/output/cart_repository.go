package output

import (
	"ddd-hexagonal-arch-go/internal/domain/entity"
	"errors"
)

var ErrNotFound = errors.New("entity not found")

// CartRepository is an interface that defines the output port for storing and retrieving shopping carts.
type CartRepository interface {
	FindByUserID(userID int) (*entity.Cart, error)
	Save(cart *entity.Cart) error
}
