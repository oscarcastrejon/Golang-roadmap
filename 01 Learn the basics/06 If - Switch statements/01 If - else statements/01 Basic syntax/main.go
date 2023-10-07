// https://yourbasic.org/golang/if-else-statement/

// Basic syntax

// An if statement executes one of two branches according to a
// boolean expression.

// If the expression evaluates to true, the if branch is executed,
// otherwise, if present, the else branch is executed.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	x := 10
	y := 30
	max := 20
	var min int

	if x > max {
		x = max
	}

	if x <= y {
		min = x
	} else {
		min = y
	}

	fmt.Println(x, max, y, min)
}
