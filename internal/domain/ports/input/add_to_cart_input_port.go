package input

type AddToCartInputPort interface {
	AddToCart(userID int, bookID int, quantity int) error
}
