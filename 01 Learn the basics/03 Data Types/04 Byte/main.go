// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// byte

// byte in Go is an alias for uint8 meaning it is an integer value.
// This integer value is of 8 bits and it represents one byte i.e number
// between 0-255). A single byte therefore can represent ASCII characters.
// Golang does not have any data type of ‘char’. Therefore

//     byte is used to represent the ASCII character
//     rune is used to represent all UNICODE characters
// 		which include every character that exists. We will
// 		study about rune later in this tutorial.

// Define Byte

// var rbyte byte := 'a'

// While declaring byte we have specify the type, as we have in the program
// above. If we don’t specify the type, then the default type is meant as a
// rune.

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
	var r byte = 'a'

	// Print Size.
	fmt.Printf("Size: %d\n", unsafe.Sizeof(r)) // OUTPUT: Size: 1

	// Print Type.
	fmt.Printf("Type: %s\n", reflect.TypeOf(r)) // OUTPUT: Type: uint8

	// Print Character.
	fmt.Printf("Character: %c\n", r) // OUTPUT: Character: a

	s := "abc"

	// This will the decimal value of byte.
	fmt.Println([]byte(s)) // OUTPUT: [97 98 99]
}
