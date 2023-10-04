// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// float32

// float32 uses single-precision floating point format to store values.
// Basically it is the set of all IEEE-754 32-bit floating-point numbers.
// The 32 bits are divided into – 1 bit sign, 8 bits exponent, and 23 bits
// mantissa. float 32 take half much size as float 64 and are comparatively
// faster on some machine architectures.

// Size: 32 bits or 4 bytes
// Range: 1.2E-38 to 3.4E+38′
// DefaultValue: 0.0

// When to Use:
//     If in your system memory is a bottleneck and range is less, then
// 		float32 can be used.

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
	// Declare a float32.
	var a float32 = 2

	// Size of float32 in bytes.
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 4 bytes
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a)) // OUTPUT: a's type is float32
}
