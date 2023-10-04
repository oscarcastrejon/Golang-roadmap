// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// uint16

// Size: 16 bits or 2 byte
// Range: 0 to 216 -1

// When to Use:
//     Use int16 when there it is known that the int range will be
// 		between 0 to 216 -1.  For temporary values such as loop invariants,
// 		it is still advisable to use int even though it might take more
// 		space because it is likely to be promoted to int in some operations
// 		or library calls.
//     For array values which lies between -0 to 216 -1, is a good use case
// 		for using int8.

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// Declare a uint16.
	var a uint16 = 2

	// Size of uint16 in bytes.
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 2 bytes
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a)) // OUTPUT: a's type is uint16
}
