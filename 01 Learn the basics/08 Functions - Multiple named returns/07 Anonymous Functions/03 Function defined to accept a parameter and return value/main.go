// https://www.golangprograms.com/go-language/functions.html

// Functions

// Anonymous Functions in Golang

// An anonymous function is a function that was declared without any
// named identifier to refer to it. Anonymous functions can accept
// inputs and return outputs, just as standard functions do.

// Function defined to accept a parameter and return value.

// Anonymous functions can be used for containing functionality that
// need not be named and possibly for short-term use.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)
}
