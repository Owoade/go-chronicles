package main

import (
	"fmt"
)

func main() {
	// Array to Slice conversion
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	fmt.Printf("%T\n", slice)

	// Slice to Array conversion
	slice2 := []int{1, 2, 3, 4, 5}
	arr2 := [5]int(slice2)
	fmt.Printf("%T\n", arr2)
}
