// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// int

// Size: Platform Dependent.
//     On 32 bit machines, the size of an int will be 32 bits or 4 bytes.
//     On 64 bit machines, the size of an int will be 64 bits or 8 bytes.

// Range: Again Platform dependent
//     On 32 bit machines, the range of int will be -231 to 231 -1.
//     On 64 bit machines, the range of int will be -263 to 263 -1.

// When to Use:
//	* It is a good idea to use int whenever using signed Integer other than
// 		the cases mentioned below
//		* When the machine is a 32 bit and range needed is greater than -231
//   		to 231 -1, then use int64 instead int. Note that in this case for int64,
// 			2 32-bit memory addresses to form a 64-bit number together.
//		* When the range is less then use appropriate integer type.

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	// This is computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64.
	sizeOfIntInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfIntInBits) // OUTPUT: 64 bits

	var a int
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 8 bytes
	fmt.Printf("a´s type is %s\n", reflect.TypeOf(a)) // OUTPUT: a´s type is int

	b := 2
	fmt.Printf("b´s type is %s\n", reflect.TypeOf(b)) // OUTPUT: b´s type is int
}
