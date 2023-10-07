// https://yourbasic.org/golang/for-loop/

// For-each range loop

// Looping over elements in slices, arrays, maps, channels or
// strings is often better done with a range loop.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
