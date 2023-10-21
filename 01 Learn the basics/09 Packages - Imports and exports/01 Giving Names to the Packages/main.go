// https://www.geeksforgeeks.org/packages-in-golang/

// Packages

// Giving Names to the Packages

// In Go language, when you name a package you must always follow
// the following points:

//     When you create a package the name of the package must be short
// 		and simple. For example strings, time, flag, etc. are standard
// 		library package.
//     The package name should be descriptive and unambiguous.
//     Always try to avoid choosing names that are commonly used or used
// 		for local relative variables.
//     The name of the package generally in the singular form. Sometimes
// 		some packages named in plural form like strings, bytes, buffers,
// 		etc. Because to avoid conflicts with the keywords.
//     Always avoid package names that already have other connotations.
// 		For example:

// =====================================================================
// To run it:
// go run main.go

// Go program to illustrate the
// concept of packages
// Package declaration
package main

// Importing multiple packages
import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	// Creating and initializing slice
	// Using shorthand declaration
	slice_1 := []byte{'*', 'G', 'e', 'e', 'k', 's', 'f', 'o', 'r', 'G', 'e', 'e', 'k', 's', '^', '^'}
	slice_2 := []string{"Gee", "ks", "for", "Gee", "ks"}

	// Displaying slices
	fmt.Println("Original Slice:")
	fmt.Printf("Slice 1 : %s", slice_1)
	fmt.Println("\nSlice 2: ", slice_2)

	// Trimming specified leading
	// and trailing Unicode points
	// from the given slice of bytes
	// Using Trim function
	res := bytes.Trim(slice_1, "*^")

	fmt.Printf("\nNew Slice : %s", res)

	// Sorting slice 2
	// Using Strings function
	sort.Strings(slice_2)

	fmt.Println("\nSorted slice:", slice_2)
}
