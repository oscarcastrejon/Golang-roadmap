// https://yourbasic.org/golang/for-loop/

// Infinite loop

// If you skip the condition as well, you get an infinite loop.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++ // Repeated forever.
	}

	// Never reached.
	fmt.Println(sum)
}
