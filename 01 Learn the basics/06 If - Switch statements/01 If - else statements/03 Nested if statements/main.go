// https://yourbasic.org/golang/if-else-statement/

// Nested if statements

// Complicated conditionals are often best expressed in Go with a
// switch statement. See 5 switch statement patterns for details.

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
	z := 20

	if x := f(); x < y {
		fmt.Println(x)
	} else if x > z {
		fmt.Println(z)
	} else {
		fmt.Println(y)
	}
}
