package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id)

	goQuote := quote.Go()
	fmt.Printf("Go saying: %s\n", goQuote)
}
