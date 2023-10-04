// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// uint32

// Size: 32 bits or 4 byte.
// Range: 0 to 232 -1.

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
	// Declare a uint32.
	var a uint32 = 2

	// Size of uint32 in bytes.
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 4 bytes
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a)) // OUTPUT: a's type is uint32
}
