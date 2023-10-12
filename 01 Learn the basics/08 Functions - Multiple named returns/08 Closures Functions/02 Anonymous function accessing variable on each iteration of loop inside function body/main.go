// https://www.golangprograms.com/go-language/functions.html

// Functions

// Closures Functions in Golang

// Closures are a special case of anonymous functions. Closures are
// anonymous functions which access the variables defined outside the
// body of the function.

// Anonymous function accessing variable on each iteration of loop
// inside function body.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()

		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)

		// fmt.Printf("%.2f Meter = %.2f Inch\n", i, func() float64 {
		// 	return i * 39.370
		// }())
	}
}
