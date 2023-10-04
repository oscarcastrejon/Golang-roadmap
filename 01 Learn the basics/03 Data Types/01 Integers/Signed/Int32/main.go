// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// int32

// Size: 32 bits or 4 byte
// Range: -231 to 231 -1.

// When to Use:

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
	// Declare a int32.
	var a int32 = 2

	// Size of int32 in bytes.
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 4 bytes
	fmt.Printf("a´s type is %s\n", reflect.TypeOf(a)) // OUTPUT: a´s type is int32
}
