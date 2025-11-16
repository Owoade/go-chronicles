package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"owoade", "anuoluwapo", "ayoade"}
	slices.Sort(strs)

	ints := []int{123, 100, 203, 458, 92}
	slices.Sort(ints)
	fmt.Println(strs, ints, slices.IsSorted(strs), slices.IsSorted(ints))
}
