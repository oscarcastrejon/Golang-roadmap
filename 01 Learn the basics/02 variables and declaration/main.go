// 1. https://go.dev/tour/basics/8
// 2. https://gobyexample.com/variables
// 3. https://www.w3schools.com/go/go_variables.php

// Variables

// The var statement declares a list of variables; as in function
// argument lists, the type is last.

// A var statement can be at package or function level. We see
// both in this example.

// =====================================================================
// Types of declarations.

// 	1. With the var keyword:
// Use the var keyword, followed by variable name and type:
// var variablename type = value

// 	2. With the := sign:
// Use the := sign, followed by the variable value:
// variablename := value

// =====================================================================
// Difference Between var and :=
// There are some small differences between the var and :=

// var
// Can be used inside and outside of functions.
// Variable declaration and value assignment can be done separately.

// :=
// Can only be used inside functions.
// Variable declaration and value assignment cannot be done separately
// (must be done in the same line).

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

// Variables 1
var c, python, java bool

func variables1() {
	var i int

	fmt.Println("V1: ", i, c, python, java) // OUTPUT: V1:  0 false false false
}

func variables2() {
	var a = "initial"
	fmt.Println("V2: ", a) // OUTPUT: V2:  initial

	var b, c int = 1, 2
	fmt.Println("V2: ", b, c) // OUTPUT: V2:  1 2

	var d = true
	fmt.Println("V2: ", d) // OUTPUT: V2:  true

	var e int
	fmt.Println("V2: ", e) // OUTPUT: V2:  0

	f := "apple"
	fmt.Println("V2: ", f) // OUTPUT: V2:  apple
}

func variables3() {
	// Variable Declaration With Initial Value.
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println("V3: ", student1) // OUTPUT: V3:  John
	fmt.Println("V3: ", student2) // OUTPUT: V3:  Jane
	fmt.Println("V3: ", x)        // OUTPUT: V3:  2

	// Variable Declaration Without Initial Value.
	var a string
	var b int
	var c bool

	fmt.Println("V3: ", a) // OUTPUT: V3:
	fmt.Println("V3: ", b) // OUTPUT: V3:  0
	fmt.Println("V3: ", c) // OUTPUT: V3:  false

	// Value Assignment After Declaration.
	var student3 string
	student3 = "John"

	fmt.Println("V3: ", student3) // OUTPUT: V3:  John
}

func main() {
	variables1()
	variables2()
	variables3()
}
