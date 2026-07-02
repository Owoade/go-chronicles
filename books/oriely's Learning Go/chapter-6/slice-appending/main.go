package main

import "fmt"

func main() {
	sl := make([]int, 10, 20)
	fmt.Println(sl) // [0 0 0 0 0 0 0 0 0 0]

	AppendSlice(sl, 20)

	fmt.Println(sl) // [0 0 0 0 0 0 0 0 0 0]
	// sl's header (len=10) was passed by value and is unaffected by AppendSlice.
}

func AppendSlice(sl []int, val int) {
	// sl is a copy of main's slice header, pointing to the same backing array.
	newSlice := append(sl, val)
	// cap (20) > len (10), so no reallocation: val is written into index 10
	// of the shared array, and a new header {len: 11} is returned as newSlice.
	// sl's own header still has len=10, so it doesn't see the new element.

	fmt.Println(sl) // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(newSlice)
}
