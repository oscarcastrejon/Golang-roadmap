// https://go.dev/blog/defer-panic-and-recover

// Panic - Recover - Defer

// Panic is a built-in function that stops the ordinary flow of control
// and begins panicking. When the function F calls panic, execution of
// F stops, any deferred functions in F are executed normally, and then
// F returns to its caller. To the caller, F then behaves like a call
// to panic. The process continues up the stack until all functions in
// the current goroutine have returned, at which point the program
// crashes. Panics can be initiated by invoking panic directly. They can
// also be caused by runtime errors, such as out-of-bounds array accesses.

// Recover is a built-in function that regains control of a panicking
// goroutine. Recover is only useful inside deferred functions. During
// normal execution, a call to recover will return nil and have no other
// effect. If the current goroutine is panicking, a call to recover will
// capture the value given to panic and resume normal execution.

// Here’s an example program that demonstrates the mechanics of panic
// and defer: (code)

// The function g takes the int i, and panics if i is greater than 3,
// or else it calls itself with the argument i+1. The function f defers
// a function that calls recover and prints the recovered value (if it
// is non-nil). Try to picture what the output of this program might
// be before reading on.

// The program will output:
// 		Calling g.
// 		Printing in g 0
// 		Printing in g 1
// 		Printing in g 2
// 		Printing in g 3
// 		Panicking!
// 		Defer in g 3
// 		Defer in g 2
// 		Defer in g 1
// 		Defer in g 0
// 		Recovered in f 4
// 		Returned normally from f.

// If we remove the deferred function from f the panic is not recovered
// and reaches the top of the goroutine’s call stack, terminating the
// program. This modified program will output:
// 		Calling g.
// 		Printing in g 0
// 		Printing in g 1
// 		Printing in g 2
// 		Printing in g 3
// 		Panicking!
// 		Defer in g 3
// 		Defer in g 2
// 		Defer in g 1
// 		Defer in g 0
// 		panic: 4
// 		panic PC=0x2a9cd8
// 		[stack trace omitted]

// For a real-world example of panic and recover, see the json package
// (https://go.dev/pkg/encoding/json/)
// from the Go standard library. It encodes an interface with a set of
// recursive functions. If an error occurs when traversing the value,
// panic is called to unwind the stack to the top-level function call,
// which recovers from the panic and returns an appropriate error value
// (see the ’error’ and ‘marshal’ methods of the encodeState type in
// encode.go (https://go.dev/src/pkg/encoding/json/encode.go)).

// The convention in the Go libraries is that even when a package uses
// panic internally, its external API still presents explicit error
// return values.

// Other uses of defer (beyond the file.Close example given earlier)
// include releasing a mutex:

// mu.Lock()
// defer mu.Unlock()

// printing a footer:

// printHeader()
// defer printFooter()

// and more.

// In summary, the defer statement (with or without panic and recover)
// provides an unusual and powerful mechanism for control flow. It can be
// used to model a number of features implemented by special-purpose
// structures in other programming languages. Try it out.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	f()

	fmt.Println("Returned normally from f.")
}

func f() {
	// To see panic, comment this defer function.
	defer func() {
		if r := recover(); r != nil {
			// "r" is the value of the panic.
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling g.")

	g(0)

	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 4 {
		fmt.Println("Panicking!")

		panic(fmt.Sprintf("%v", i))
		// panic(fmt.Sprint("Hola"))
	}

	defer fmt.Println("Defer in g", i)

	fmt.Println("Printing in g", i)

	g(i + 1)
}
