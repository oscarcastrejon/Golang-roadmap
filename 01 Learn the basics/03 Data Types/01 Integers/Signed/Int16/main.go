// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// int16

// Size: 16 bits or 2 byte.
// Range: -215 to 215 -1.

// When to Use:
//     Use int16 when there it is known that the int range will be between -215 to 215 -1.
// 		For temporary values such as loop invariants, it is still advisable to use int even
// 		though it might take more space because it is likely to be promoted to int in some
// 		operations or library calls.
//     For array values which lies between -215 to 215 -1, is a good use case for using int8.
// 		For eg if you are storing ASCII index for lowercase letters than int16 can be used.

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
	// Declare a int16.
	var a int16 = 2

	// Size of int16 in bytes.
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 2 bytes
	fmt.Printf("a´s type is %s\n", reflect.TypeOf(a)) // OUTPUT: a´s type is int16
}
