package main

import (
	"fmt"
)

var m = make(map[int][]int)
var sums []int

func main() {
	// arr := []int{3, 4, 2}
	// w := maxSumArray(arr, 3)
	// fmt.Println("max sum array", w)
	fmt.Println(longestUniqueSubstring("cdinedos"))
}
func maxSumArray(a []int, k int) int {
	if len(a) <= k {
		return sumArray(a)
	}
	var i int = 0

	for (i + k) <= len(a) {
		sum := sumArray(a[i : i+k])
		sums = append(sums, sum)
		m[sum] = a[i : i+k]
		i++
	}

	if len(sums) == 0 {
		return maxInSlice(a)
	}

	return maxInSlice(sums)
}

func sumArray(a []int) int {
	var sum int
	for _, val := range a {
		sum += val
	}
	return sum
}

func maxInSlice(a []int) int {
	max := a[0]
	for _, val := range a {
		if val > max {
			max = val
		}
	}
	return max
}
