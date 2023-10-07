// https://yourbasic.org/golang/if-else-statement/

// With init statement

// An if statement executes one of two branches according to a
// boolean expression.

// If the expression evaluates to true, the if branch is executed,
// otherwise, if present, the else branch is executed.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func f() int {
	return 10
}

func main() {
	y := 30

	if x := f(); x <= y {
		fmt.Println(x, y)
	}
}
