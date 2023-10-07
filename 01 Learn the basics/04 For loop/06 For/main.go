// https://gobyexample.com/for

// For

// for is Go’s only looping construct. Here are some basic types
// of for loops.

// The most basic type, with a single condition.

// A classic initial/condition/after for loop.

// for without a condition will loop repeatedly until you break out
// of the loop or return from the enclosing function.

// You can also continue to the next iteration of the loop.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i) // OUTPUT: 1 \n 2 \n 3

		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j) // OUTPUT: 7 \n 8 \n 9
	}

	for {
		fmt.Println("loop") // OUTPUT: loop

		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n) // OUTPUT: 1 \n 3 \n 5
	}
}
