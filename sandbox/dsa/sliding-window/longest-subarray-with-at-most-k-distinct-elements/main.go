package main

import (
	"fmt"
	"slices"
)

func main() {
	println(longestSubarrayWithKMostDistinctElement(
		[]int{1, 2, 3, 5, 6},
		3,
	))
}

func longestSubarrayWithKMostDistinctElement(a []int, k int) int {
	var i = 0
	var j = 1
	if len(a) < k {
		return len(a)
	}
	w := []int{a[0]}
	var res = []int{}
	var lastEntry bool
	tolerance := 1
	for i < len(a) && !lastEntry {
		nextEl := a[j]
		if slices.Contains(w, nextEl) {
			lastEntry = j == (len(a) - 1)
			w = append(w, nextEl)
			if !slices.Contains(res, len(w)) {
				res = append(res, len(w))
			}
			if (j + 1) < len(a) {
				j++
			}
		} else if tolerance < k {
			lastEntry = j == (len(a) - 1)
			w = append(w, nextEl)
			if !slices.Contains(res, len(w)) {
				res = append(res, len(w))
			}
			tolerance++
			if (j + 1) < len(a) {
				j++
			}
		} else {
			firstElInWindow := w[0]
			if len(a) > 1 {
				w = w[1:]
				if !slices.Contains(w, firstElInWindow) {
					tolerance--
				}
			} else {
				w = []int{}
			}
			i++
		}
		fmt.Println(i, j, w, res, lastEntry)
	}

	return maxInSlice(res)
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
