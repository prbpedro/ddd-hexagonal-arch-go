package main

import (
	"ddd-hexagonal-arch-go/internal/application/initializer"
	"fmt"
)

func main() {

	applicationInitializer := initializer.NewInMemoryDatabaseApplicationInitializer()

	searchBooksUseCase, addCartUseCase, err := applicationInitializer.InitApp()
	panicIfErr(err)

	// Search for books
	books, err := searchBooksUseCase.Search("Book 1")
	panicIfErr(err)
	fmt.Printf("Found %d books\n", len(books))

	// Add books to cart
	err = addCartUseCase.AddToCart(1, 1, 2)
	panicIfErr(err)

	fmt.Println("Book added to cart")
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
