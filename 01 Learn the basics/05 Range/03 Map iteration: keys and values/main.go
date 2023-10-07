// https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/

// Map iteration: keys and values

// The iteration order over maps is not specified and is not guaranteed to
// 	be the same from one iteration to the next.

// If a map entry that has not yet been reached is removed during
// 	iteration, this value will not be produced.
// If a map entry is created during iteration, that entry may or may
// 	not be produced.
// For a nil map, the number of iterations is 0.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range m {
		fmt.Println(k, v)
		/*	OUTPUT:
			two 2
			three 3
			one 1
		*/
	}
}
