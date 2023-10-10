// https://www.golangprograms.com/go-language/functions.html

// Functions

// Golang Functions Returning Multiple Values

// Functions in Golang can return multiple values, which is a helpful
// feature in many practical scenarios.

// This example declares a function with two return values and calls
// it from a main function.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b

	// Return statement without specify variable name.
	return
}

func main() {
	var a, p int

	a, p = rectangle(20, 30)

	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
}
