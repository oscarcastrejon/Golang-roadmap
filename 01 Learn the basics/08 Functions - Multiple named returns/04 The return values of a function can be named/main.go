// https://www.golangprograms.com/go-language/functions.html

// Functions

// The return values of a function can be named in Golang

// Golang allows you to name the return values of a function. We can
// also name the return value by defining variables, here a variable
// total of integer type is defined in the function declaration for
// the value that the function returns.

// Since the function is declared to return a value of type int, the
// last logical statement in the execution flow must be a return
// statement that returns a value of the declared type.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func rectangle(l int, b int) (area int) {
	var parameter int

	parameter = 2 * (l + b)

	fmt.Println("Parameter: ", parameter)

	area = l * b

	// Return statement without specify variable name.
	return
}

func main() {
	fmt.Println("Area: ", rectangle(20, 30))
}
