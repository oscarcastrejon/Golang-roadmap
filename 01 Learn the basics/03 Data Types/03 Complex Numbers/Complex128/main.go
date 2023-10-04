// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// complex128

// For complex128 both real and imaginary part are float64

// Size: Both real and imaginary part are of same size as float64.
// It is of size 64 bits or 8 bytes

// Range: Both real and imaginary part range is same as float64
// i.e -1.7E+308 to +1.7E+308

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
	var a float64 = 3
	var b float64 = 5

	fmt.Printf("a's size is %d bytes\n", unsafe.Sizeof(a)) // OUTPUT: a's size is 8 bytes
	fmt.Printf("b's size is %d bytes\n", unsafe.Sizeof(b)) // OUTPUT: b's size is 8 bytes

	// Initialize -1.
	c := complex(a, b)

	fmt.Println(c) // OUTPUT: (3+5i)

	// Initialize -2. When don't specify a type , the default type will be complex128.
	d := 4 + 5i

	// Print size.
	fmt.Printf("c's size is %d bytes\n", unsafe.Sizeof(c)) // OUTPUT: c's size is 16 bytes
	fmt.Printf("d's size is %d bytes\n", unsafe.Sizeof(d)) // OUTPUT: d's size is 16 bytes

	// Print type.
	fmt.Printf("c's type is %s\n", reflect.TypeOf(c)) // OUTPUT: c's type is complex128
	fmt.Printf("d's type is %s\n", reflect.TypeOf(d)) // OUTPUT: d's type is complex128

	//Operations on complex number.
	fmt.Println(c+d, c-d, c*d, c/d) // OUTPUT: (7+10i) (-1+0i) (-13+35i) (0.902439024390244+0.12195121951219513i)
}
