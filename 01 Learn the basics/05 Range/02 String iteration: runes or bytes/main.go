// https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/

// String iteration: runes or bytes

// The index is the first byte of a UTF-8-encoded code point; the second
// 	value, of type rune, is the value of the code point.
// For an invalid UTF-8 sequence, the second value will be 0xFFFD, and
// 	the iteration will advance a single byte.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	for i, ch := range "日本語" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
		/* OUTPUT:
		U+65E5 '日' starts at byte position 0
		U+672C '本' starts at byte position 3
		U+8A9E '語' starts at byte position 6
		*/
	}

	// To loop over individual bytes, simply use a normal for loop and
	// string indexing:
	/*
		const s = "日本語"
		for i := 0; i < len(s); i++ {
			fmt.Printf("%x ", s[i]) // OUTPUT: e6 97 a5 e6 9c ac e8 aa 9e
		}
	*/
}
