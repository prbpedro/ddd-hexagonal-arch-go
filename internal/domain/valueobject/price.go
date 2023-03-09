package valueobject

// Price is a struct that represents the price of a book. It has a Value field of type float64.
type Price struct {
	value float64
}

func NewPrice(value float64) *Price {
	return &Price{
		value: value,
	}
}

func (p *Price) Value() float64 {
	return p.value
}
