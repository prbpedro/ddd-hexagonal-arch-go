package entity

// Cart is a struct that represents a shopping cart in the domain model.
type Cart struct {
	UserID int
	Items  []*CartItem
}

func NewCart(userID int, items []*CartItem) *Cart {
	return &Cart{
		UserID: userID,
		Items:  items,
	}
}

// CartItem is a struct that represents an item in a shopping cart.
type CartItem struct {
	BookID   int
	Quantity int
}

func NewCartItem(bookID int, quantity int) *CartItem {
	return &CartItem{
		BookID:   bookID,
		Quantity: quantity,
	}
}

// AddItem adds an item to the shopping cart.
func (c *Cart) AddItem(item *CartItem) {
	for _, i := range c.Items {
		if i.BookID == item.BookID {
			i.Quantity += item.Quantity
			return
		}
	}
	c.Items = append(c.Items, item)
}
