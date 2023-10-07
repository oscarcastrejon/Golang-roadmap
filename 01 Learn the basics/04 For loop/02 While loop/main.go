// https://yourbasic.org/golang/for-loop/

// While loop

// If you skip the init and post statements, you get a while loop.

// The condition, n < 5, is computed.

// If true, the loop body runs,
// otherwise the loop is done.

// Back to step 1.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	n := 1
	for n < 5 {
		n *= 2
	}

	fmt.Println(n) // OUTPUT: 8
}
