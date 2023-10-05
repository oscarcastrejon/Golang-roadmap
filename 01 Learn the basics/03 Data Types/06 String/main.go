// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// String
// string is a read only slice of bytes in golang. String can be
// initialized in two ways

//     using double quotes “” eg “this”

// string in double quotes honors the escape sequences. For eg if
// the string contains a \n then while printing there will be a new
// line

//     using back quotes ` eg  \`this`

// String in back quotes is just a raw string and it does not honor
// any kind of escape sequences.

// Each character in a string will occupy some bytes depending upon
// encoding used. For eg in utf-8 encoded string, each character will
// occupy between 1-4 bytes. You can read about utf-8 in this must
// read famous blog-The Absolute Minimum Every Software Developer Absolutely,
// Positively Must Know About Unicode and Character Sets (No Excuses!).
// In utf-8 , the characters a or b are encoded using 1  byte while the
// character pound sign £ is encoded using two bytes . Therefore the string
// “ab£” will output 4 bytes when you will convert the string to byte array
// and print it like below

// s := "ab£"
// fmt.Println([]byte(s))

// Output
// [48 98 194 163]

// Also when you try to print the length of the above string using len(“ab£”),
// it will output 4 and not 3 because it contains 4 bytes.

// Also note that range loops over sequences of byte which form each character,
// therefore for the below range loop

// for _, c := range s {
//    fmt.Println(string(c))
// }

// Output will be
// a
// b
// £

// There are many operations that can be performed on a string. One such
// operation is concatenation which combines two string. The sign ‘+’ is
// used for concatenation. Let’s see full working  code for all above
// things that we discussed

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"fmt"
)

func main() {
	// String in double quotes.
	x := "this\nthat"

	// OUTPUT:
	// x is: this
	// that
	fmt.Printf("x is: %s\n", x)

	// String in back quotes.
	y := `this\nthat`

	fmt.Printf("y is: %s\n", y) // OUTPUT: y is: this\nthat

	s := "ab£"

	// This will print the byte sequence.
	// Since character a and b occupies 1 byte each and £ character occupies 2 bytes.
	// The final output will 4 bytes.
	fmt.Println([]byte(s)) // OUTPUT: [97 98 194 163]

	// The output will be 4 for same reason as above.
	fmt.Println(len(s)) // OUTPUT: 4

	// range loops over sequences of byte which form each character.
	for _, c := range s {
		fmt.Println(string(c)) // OUTPUT: a; b; £
	}

	// Concatenation.
	fmt.Println("c" + "d") // OUTPUT: cd
}
