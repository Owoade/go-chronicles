package main

import (
	"fmt"
	"testing"
)

func TestMaxSumWindow(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
		desc     string
	}{
		// ✅ Normal cases
		{[]int{2, 1, 5, 1, 3, 2}, 3, 9, "window [5,1,3] in middle"},
		{[]int{1, 2, 3, 4, 5}, 2, 9, "window [4,5] at end"},
		{[]int{5, 1, 1, 1, 5}, 2, 6, "max window at edges [5,1] vs [1,5]"},

		// 🔲 Window = full slice
		{[]int{3, 4, 2}, 3, 9, "k equals slice length"},

		// 🔢 Single element window
		{[]int{10, 2, 3}, 1, 10, "k=1, max single element"},

		// 📉 All same elements
		{[]int{4, 4, 4, 4}, 2, 8, "all same, any window of 2"},

		// 📈 Increasing order
		{[]int{1, 2, 3, 4, 5}, 3, 12, "increasing, window [3,4,5]"},

		// 📉 Decreasing order
		{[]int{5, 4, 3, 2, 1}, 3, 12, "decreasing, window [5,4,3]"},

		// 🔴 Negative numbers
		{[]int{-1, -2, -3, -4}, 2, -3, "all negative, window [-1,-2]"},
		{[]int{-5, 3, -2, 7}, 2, 5, "mixed neg/pos, window [-2,7]"},

		// 1️⃣ Single element slice
		{[]int{42}, 1, 42, "single element slice"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			result := maxSumArray(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("FAIL [%s]\n  nums=%v k=%d\n  got=%d want=%d",
					tt.desc, tt.nums, tt.k, result, tt.expected)
			} else {
				fmt.Printf("PASS [%s] → %d\n", tt.desc, result)
			}
		})
	}
}