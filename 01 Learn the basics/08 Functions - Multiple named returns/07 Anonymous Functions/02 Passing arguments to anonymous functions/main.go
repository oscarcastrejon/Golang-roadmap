// https://www.golangprograms.com/go-language/functions.html

// Functions

// Anonymous Functions in Golang

// An anonymous function is a function that was declared without any
// named identifier to refer to it. Anonymous functions can accept
// inputs and return outputs, just as standard functions do.

// Passing arguments to anonymous functions.

// Anonymous functions can be used for containing functionality that
// need not be named and possibly for short-term use.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)
}
