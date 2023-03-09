package initializer

import "ddd-hexagonal-arch-go/internal/domain/ports/input"

type ApplicationInitializer interface {
	InitApp() (input.SearchBooksInputPort, input.AddToCartInputPort, error)
}
