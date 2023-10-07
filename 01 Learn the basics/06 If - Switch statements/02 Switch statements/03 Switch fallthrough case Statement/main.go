// https://www.golangprograms.com/golang-switch-case-statements.html

// Golang - switch fallthrough case Statement

// The fallthrough keyword used to force the execution flow to fall
// through the successive case block.

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
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}
