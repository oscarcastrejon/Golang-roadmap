// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// uint

// Size: Platform Dependent.
//     On 32 bit machines, the size of an int will be 32 bits or 4 byte.
//     On 64 bit machines, the size of an int will be 64 bits or 8 byte

// Range: Again Platform dependent
//     On 32 bit machines, the range of int will be -231 to 231 -1.
//     On 64 bit machines, the range of int will be -263 to 263 -1

// When to Use:
//     It is a good idea to use uint  whenever using signed Integer other
// 		than the cases mention below
//         When the machine is a 32 bit and range needed is greater than
// 			-231 to 231 -1, then use int64 instead int. Note that in this
// 			case for int64,  2 32-bit memory addresses to form a 64-bit
// 			number together.
//         When the range is less than use the appropriate int type.

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
	// This is computed as const uintSize = 32 << (^uuint(0) >> 32 & 1) // 32 or 64.
	sizeOfuintInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfuintInBits) // OUTPUT: 64 bits

	var a uint
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // OUTPUT: 8 bytes
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a)) // OUTPUT: a's type is uint
}
