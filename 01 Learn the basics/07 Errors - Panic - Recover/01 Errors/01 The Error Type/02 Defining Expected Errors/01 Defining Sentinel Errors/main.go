// https://earthly.dev/blog/golang-errors/

// Errors

// The Error Type

// Defining Expected Errors

// Defining Sentinel Errors

// Building on the Divide function from earlier, we can improve the
// error signaling by pre-defining a “Sentinel” error. Calling functions
// can explicitly check for this error using errors.Is:

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return a / b, nil
}

func main() {
	a, b := 10, 0

	result, err := Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}

		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}
