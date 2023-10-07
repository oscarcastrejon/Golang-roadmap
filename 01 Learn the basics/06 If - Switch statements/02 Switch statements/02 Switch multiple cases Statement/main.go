// https://www.golangprograms.com/golang-switch-case-statements.html

// Golang - switch multiple cases Statement

// The switch with multiple case line statement is used to select
// common block of code for many similar cases.

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
	var t int = today.Day()

	switch t {
	case 5, 10, 15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}
