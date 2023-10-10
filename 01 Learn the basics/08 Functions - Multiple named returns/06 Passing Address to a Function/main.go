// https://www.golangprograms.com/go-language/functions.html

// Functions

// Golang Passing Address to a Function

// Passing the address of variable to the function and the value of a
// variables modified using dereferencing inside body of function.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address.
	*t = *t + " Doe" // defrencing pointer address.

	return
}

func main() {
	var age = 20
	var text = "John"

	fmt.Println("Before:", text, age)

	update(&age, &text)

	fmt.Println("After :", text, age)
}
