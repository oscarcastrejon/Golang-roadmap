// https://earthly.dev/blog/golang-errors/

// Errors

// Constructing Errors

// Errors can be constructed on the fly using Go’s built-in errors
// or fmt packages. For example, the following function uses the
// errors package to return a new error with a static error message:
// (DoSomething)

// Similarly, the fmt package can be used to add dynamic data to the
// error, such as an int, string, or another error. For example:
// (Divide)

// Note that fmt.Errorf will prove extremely useful when used to wrap
// another error with the %w format verb - but I’ll get into more detail
// on that further down in the article.

// There are a few other important things to note in the example above.

// Errors can be returned as nil, and in fact, it’s the default,
// or “zero”, value of on error in Go. This is important since
// checking if err != nil is the idiomatic way to determine if
// an error was encountered (replacing the try/catch statements
// you may be familiar with in other programming languages).

// Errors are typically returned as the last argument in a function.
// Hence in our example above, we return an int and an error, in
// that order.

// When we do return an error, the other arguments returned by the
// function are typically returned as their default “zero” value.
// A user of a function may expect that if a non-nil error is
// returned, then the other arguments returned are not relevant.

// Lastly, error messages are usually written in lower-case and
// don’t end in punctuation. Exceptions can be made though, for
// example when including a proper noun, a function name that
// begins with a capital letter, etc.

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"errors"
	"fmt"
)

func DoSomething() error {
	return errors.New("something didn't work")
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}

	return a / b, nil
}

func main() {
	fmt.Println(DoSomething())
	fmt.Println(Divide(6, 0))
}
