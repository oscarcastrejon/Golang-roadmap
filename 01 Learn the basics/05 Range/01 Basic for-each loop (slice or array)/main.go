// https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/

// Basic for-each loop (slice or array)

// The range expression, a, is evaluated once before beginning the loop.
// The iteration values are assigned to the respective iteration variables,
// 		i and s, as in an assignment statement.
// The second iteration variable is optional.
// For a nil slice, the number of iterations is 0.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	a := []string{"Foo", "Bar"}
	for i, s := range a {
		fmt.Println(i, s) // OUTPUT: 0 Foo \n 1 Bar
	}
}
