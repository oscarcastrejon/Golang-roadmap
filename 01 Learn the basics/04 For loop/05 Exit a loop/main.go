// https://yourbasic.org/golang/for-loop/

// Exit a loop

// The break and continue keywords work just as they do in C and Java.

// A continue statement begins the next iteration of the innermost for loop at its post statement (i++).
// A break statement leaves the innermost for, switch or select statement.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // Skip odd numbers.
			continue
		}

		sum += i
	}

	fmt.Println(sum) // OUTPUT: 6
}
