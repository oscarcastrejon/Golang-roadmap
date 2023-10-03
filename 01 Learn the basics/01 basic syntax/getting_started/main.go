// https://go.dev/doc/tutorial/getting-started

// Call code in an external package.

// To run it:
// go run main.go

// This method to run, only works when a module was created.
// go run .

package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
