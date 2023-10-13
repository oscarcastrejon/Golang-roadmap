// https://www.golangprograms.com/go-language/functions.html

// Functions

// Higher Order Functions in Golang

// A Higher-Order function is a function that receives a function as an
// argument or returns the function as output.

// Higher order functions are functions that operate on other functions,
// either by taking them as arguments or by returning them.

// Returning Functions from other Functions

// In the program above, the squareSum function signature specifying
// that function returns two functions and one integer value.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}
func main() {
	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum(5)(6)(7))
}
