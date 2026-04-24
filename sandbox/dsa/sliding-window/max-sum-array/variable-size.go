package main

import "strings"

func longestUniqueSubstring(a string) int {
	var sums []int
	var i = 0
	var j = 1
	var cur = string(a[0])
	for i < len(a) {
		nextChar := string(a[j])
		if strings.Contains(cur, nextChar) {
			i++
			sums = append(sums, len(cur))
			cur = string(cur[1:])
		} else {
			if (j + 1) < len(a) {
				j++
			}
			cur += nextChar
		}
	}

	if len(sums) == 0 {
		return 0
	}

	return maxInSlice(sums)
}
