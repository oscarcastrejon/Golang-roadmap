// https://www.golangprograms.com/golang-switch-case-statements.html

// Golang - swith conditional cases Statement

// The case statement can also used with conditional operators.

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch {
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
