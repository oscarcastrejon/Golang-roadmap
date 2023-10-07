// https://yourbasic.org/golang/for-loop/

// Three-component loop

// This version of the Go for loop works just as in C or Java.

// The init statement, i := 1, runs.
// The condition, i < 5, is computed.

// If true, the loop body runs,
// otherwise the loop is done.

// The post statement, i++, runs.
// Back to step 2.

// The scope of i is limited to the loop.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}

	fmt.Println(sum) // OUTPUT: 10 (1+2+3+4)
}
