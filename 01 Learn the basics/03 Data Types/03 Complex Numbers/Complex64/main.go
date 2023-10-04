// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// complex64

// For complex 64 both real and imaginary part are float32.

// Size: Both real and imaginary part are of same size as float32.
// It is of size 32 bits or 4 bytes

// Range: Both real and imaginary part range is same as float32
// i.e 1.2E-38 to 3.4E+38

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
	var a float32 = 3
	var b float32 = 5

	fmt.Printf("a's size is %d bytes\n", unsafe.Sizeof(a)) // OUTPUT: a's size is 4 bytes
	fmt.Printf("b's size is %d bytes\n", unsafe.Sizeof(b)) // OUTPUT: b's size is 4 bytes

	// Initialize -1.
	c := complex(a, b)

	fmt.Println(c) // OUTPUT: (3+5i)

	// Initialize -2.
	var d complex64
	d = 4 + 5i

	// Print size.
	fmt.Printf("c's size is %d bytes\n", unsafe.Sizeof(c)) // OUTPUT: c's size is 8 bytes
	fmt.Printf("d's size is %d bytes\n", unsafe.Sizeof(d)) // OUTPUT: d's size is 8 bytes

	// Print type.
	fmt.Printf("c's type is %s\n", reflect.TypeOf(c)) // OUTPUT: c's type is complex64
	fmt.Printf("d's type is %s\n", reflect.TypeOf(d)) // OUTPUT: d's type is complex64

	//Operations on complex number.
	fmt.Println(c+d, c-d, c*d, c/d) // OUTPUT: (7+10i) (-1+0i) (-13+35i) (0.902439+0.12195122i)
}
