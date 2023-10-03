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
// Can be used inside and outside of functions
// Variable declaration and value assignment can be done separately.

// :=
// Can only be used inside functions
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

	fmt.Println(i, c, python, java)
}

func variables2() {
	var a = "initial"
	fmt.Println("V2: ", a)

	var b, c int = 1, 2
	fmt.Println("V2: ", b, c)

	var d = true
	fmt.Println("V2: ", d)

	var e int
	fmt.Println("V2: ", e)

	f := "apple"
	fmt.Println("V2: ", f)
}

func variables3() {
	// Variable Declaration With Initial Value.
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println("V3: ", student1)
	fmt.Println("V3: ", student2)
	fmt.Println("V3: ", x)

	// Variable Declaration Without Initial Value.
	var a string
	var b int
	var c bool

	fmt.Println("V3: ", a)
	fmt.Println("V3: ", b)
	fmt.Println("V3: ", c)

	// Value Assignment After Declaration
	var student3 string
	student3 = "John"

	fmt.Println("V3: ", student3)
}

func main() {
	variables1()
	variables2()
	variables3()
}
