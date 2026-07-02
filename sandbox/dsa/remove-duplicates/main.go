package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	nums = removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) []int {
	hash := make(map[int][]int)
	i := 0
	for i < len(nums) {
		el := nums[i]
		hash[el] = append(hash[el], el)
		copiedArr := make([]int, len(nums))
		copy(copiedArr, nums)
		if slices.Contains(nums[:i], el) {
			if (i + 1) < len(nums) {
				fmt.Println(i, nums, copiedArr[i+1:], "already exists: ", true)
				copy(nums[i:], copiedArr[i+1:])
			}
		} else {
			fmt.Println(i, nums, copiedArr[i:], "already exists: ", false)
			copy(nums[i:], copiedArr[i:])
		}
		fmt.Println("	merged result: ", nums)
		i = i + 1
	}
	return nums
}
