// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// uint64

// Size: 64 bits or 8 byte.
// Range: 0 to 264 -1.

// When to Use:
//     uint64 is used when the range is higher.

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
	// Declare a uint64.
	var a uint64 = 2

	// Size of uint64 in bytes.
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 8 bytes
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a)) // OUTPUT: a's type is uint64
}
