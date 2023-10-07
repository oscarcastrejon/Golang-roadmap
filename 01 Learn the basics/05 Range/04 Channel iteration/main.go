// https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/

// Map iteration: keys and values

// For channels, the iteration values are the successive values
// 	sent on the channel until closed.

// For a nil channel, the range loop blocks forever.

// =====================================================================
// To run it:
// go run main.go

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
		/*	OUTPUT:
			1
			2
			3
		*/
	}
}
