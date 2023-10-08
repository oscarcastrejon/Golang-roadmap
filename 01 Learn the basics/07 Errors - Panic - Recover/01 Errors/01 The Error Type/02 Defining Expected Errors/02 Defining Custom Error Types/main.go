// https://earthly.dev/blog/golang-errors/

// Errors

// The Error Type

// Defining Expected Errors

// Defining Custom Error Types

// Many error-handling use cases can be covered using the strategy
// above, however, there can be times when you might want a little
// more functionality. Perhaps you want an error to carry additional
// data fields, or maybe the error’s message should populate itself
// with dynamic values when it’s printed.

// You can do that in Go by implementing custom errors type.

// Below is a slight rework of the previous example. Notice the new
// type DivisionError, which implements the Error interface. We can
// make use of errors.As to check and convert from a standard error
// to our more specific DivisionError.

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"errors"
	"fmt"
)

type DivisionError struct {
	IntA int
	IntB int
	Msg  string
}

func (e *DivisionError) Error() string {
	return e.Msg
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{
			IntA: a,
			IntB: b,
			Msg:  fmt.Sprintf("cannot divide '%d' by zero", a),
		}
	}

	return a / b, nil
}

func main() {
	a, b := 10, 0

	result, err := Divide(a, b)
	if err != nil {
		var divErr *DivisionError

		switch {
		case errors.As(err, &divErr):
			fmt.Printf("%d / %d is not mathematically valid: %s\n", divErr.IntA, divErr.IntB, divErr.Error())
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}

		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}
