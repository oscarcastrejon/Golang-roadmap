// int8

// Size: 8 bits or 1 byte
// Range: -27 to 27 -1.

// When to Use:
//     Use int8 when there it is known that the int range will be
//		between -27 to 27 -1.  For temporary values such as loop
// 		invariants, it is still advisable to use int even though
// 		it might take more space because it is likely to be promoted
// 		to int in some operations or library calls.
//     For array values which lies between -27 to 27 -1, is a good
// 		use case for using int8. For eg if you are storing ASCII index
// 		for lowercase letters then int8 can be used
//     It is a good idea to use int8 for data values.

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a int 8
	var a int8 = 2

	//Size of int8 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))        // 1 bytes
	fmt.Printf("a´s type is %s\n", reflect.TypeOf(a)) // a´s type is int8
}
