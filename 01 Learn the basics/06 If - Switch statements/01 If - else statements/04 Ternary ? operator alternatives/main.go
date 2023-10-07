// https://yourbasic.org/golang/if-else-statement/

// Ternary ? operator alternatives

// You can’t write a short one-line conditional in Go; there is
// no ternary conditional operator. Instead of:

// res = expr ? x : y

// you write

// if expr {
// 	res = x
// } else {
// 	res = y
// }

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func Min(x, y int) int {
	if x <= y {
		return x
	}

	return y
}

func main() {
	fmt.Println(Min(10, 20))
}
