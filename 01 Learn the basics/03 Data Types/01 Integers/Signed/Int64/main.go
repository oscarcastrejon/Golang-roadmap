// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// int64

// Size: 64 bits or 8 byte
// Range: -263 to 263 -1

// When to Use:
//     int64 is used when the range is higher. For eg time.Duration is of type int64

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
	// Declare a int64.
	var a int64 = 2

	// Size of int64 in bytes.
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 8 bytes
	fmt.Printf("a´s type is %s\n", reflect.TypeOf(a)) // OUTPUT: a´s type is int64
}
