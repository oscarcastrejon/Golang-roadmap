// https://golangbyexample.com/all-data-types-in-golang-with-examples/

// Booleans

// The data type is bool and has two possible values true or false.

// Default Value: false

// Operations:
//     AND – &&
//     OR  – ||
//     Negation – !

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	//Default value will be false it not initialized
	var a bool
	fmt.Printf("a's value is %t\n", a) // OUTPUT: a's value is false

	//And operation on one true and other false
	andOperation := 1 < 2 && 1 > 3
	fmt.Printf("Ouput of AND operation on one true and other false %t\n", andOperation) // OUTPUT: Ouput of AND operation on one true and other false false

	//OR operation on one true and other false
	orOperation := 1 < 2 || 1 > 3
	fmt.Printf("Ouput of OR operation on one true and other false: %t\n", orOperation) // OUTPUT: Ouput of OR operation on one true and other false: true

	//Negation Operation on a false value
	negationOperation := !(1 > 2)
	fmt.Printf("Ouput of NEGATION operation on false value: %t\n", negationOperation) // OUTPUT: Ouput of NEGATION operation on false value: true
}
