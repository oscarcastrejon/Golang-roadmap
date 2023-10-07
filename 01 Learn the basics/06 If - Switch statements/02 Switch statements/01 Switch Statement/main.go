// https://www.golangprograms.com/golang-switch-case-statements.html

// Golang - switch Statement

// The switch statement is used to select one of many blocks of
// code to be executed.

// Consider the following example, which display a different
// message for particular day.

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

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}
