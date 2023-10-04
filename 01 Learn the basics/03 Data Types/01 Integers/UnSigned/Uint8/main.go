// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// uint8

// Size: 8 bits or 1 byte.
// Range:  0 to 255 or 0 to 28 -1.

// When to Use:
//     Use uint8 when there it is known that the int range will be between 28 -1.
// 		For temporary values such as loop invariants, it is still advisable to use
// 		int even though it might take more space because it is likely to be promoted
// 		to int in some operations or library calls.
//     For array values which lies between  28 -1. is a good use case for using uint8.
// 		For eg if you are storing ASCII index in an array then uint8 can be used.

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
	// Declare a uint8.
	var a uint8 = 2

	// Size of uint8 in bytes.
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 1 bytes
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a)) // OUTPUT: a's type is uint8
}
