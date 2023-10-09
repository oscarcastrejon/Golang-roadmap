// https://www.golangprograms.com/go-language/functions.html

// Functions

// Creating a Function in Golang

// A declaration begins with the func keyword, followed by the name
// you want the function to have, a pair of parentheses (), and then a
// block containing the function's code.

// The following example has a function with the name SimpleFunction.
// It takes no parameter and returns no values.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

// SimpleFunction prints a message.
func SimpleFunction() {
	fmt.Println("Hello World")
}

func main() {
	SimpleFunction()
}
