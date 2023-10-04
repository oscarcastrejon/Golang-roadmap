// // https://golangbyexample.com/all-data-types-in-golang-with-examples/

// uintptr

// This is an unsigned integer type that is large enough to hold any
// pointer address. Therefore is size and range are platform dependent.

// Size: Platform Dependent.
//     On 32 bit machines, the size of int will be 32 bits or 4 byte.
//     On 64 bit machines, the size of int will be 64 bits or 8 byte

// Range: Again Platform dependent
//     On 32 bit machines, the range of int will be -231 to 231 -1.
//     On 64 bit machines, the range of int will be -263 to 263 -1

// Properties:
//     A uintptr can be converted to unsafe.Pointer and viceversa
//     Arithmetic can be performed on the uintptr
//     uintptr even though it holds a pointer address, is just a value
// 		and does not references any object. Therefore
//         Its value will not be updated if the corresponding object moves.
// 			Eg When goroutine stack changes
//         The corresponding object can be garbage collected.

// When to Use:
//     Its purpose is to be used along with unsafe.Pointer mainly used for
// 		unsafe memory access.
//     When you want to save the pointer address value for printing it or
// 		storing it. Since the address is just stored and does not reference
// 		anything, the corresponding object can be garbage collected.

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"fmt"
	"unsafe"
)

type sample struct {
	a int
	b string
}

func main() {
	s := &sample{a: 1, b: "test"}

	// Getting the address of field b in struct s.
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))

	// Typecasting it to a string pointer and printing the value of it.
	fmt.Println(*(*string)(p)) // OUTPUT: test
}
