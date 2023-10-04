// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// float64

// float64 uses a double-precision floating-point format to store values.
// Basically it is the set of all IEEE-754 64-bit floating-point numbers.
// The 64 bits are divided into – 1-bit sign, 11 bits exponent, 52 bits
// mantissa. float64 takes twice as much size compared to float32 but can
// represent numbers more accurately than float32.

// Size: 32 bits or 4 bytes
// Range: 1.2E-38 to 3.4E+38
// DefaultValue: 0.0

// When to Use:
//     When the precision needed is high.

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
	// Declare a float64.
	var a float64 = 2

	// Size of float64 in bytes.
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 8 bytes
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a)) // OUTPUT: a's type is float64

	// Default is float64 when you don't specify a type.
	b := 2.3
	fmt.Printf("b's type is %s\n", reflect.TypeOf(b)) // OUTPUT: b's type is float64
}
