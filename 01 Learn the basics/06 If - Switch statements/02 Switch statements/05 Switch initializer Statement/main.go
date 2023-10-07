// https://www.golangprograms.com/golang-switch-case-statements.html

// Golang - switch initializer Statement

// The switch keyword may be immediately followed by a simple
// initialization statement where variables, local to the switch
// code block, may be declared and initialized.

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"fmt"
	"time"
)

func main() {
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}

}
