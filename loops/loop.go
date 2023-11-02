package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 8}

	fmt.Println(get_max_element(arr))

	test_arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(remove_duplicate((test_arr)))

}

func get_max_element(arr []int) (max int) {

	if len(arr) == 0 {
		return -1
	}

	_max := arr[0]

	for i := 0; i < len(arr); i++ {

		if arr[i] > max {
			_max = arr[i]
		}
	}

	return _max

}

func remove_duplicate(arr []int) ([]int, []string) {

	var duplicate []string

	var new_arr []int

	for i, el := range arr {

		if i == 0 {
			new_arr = append(new_arr, el)
			continue
		}

		prev_el := arr[i-1]

		if el == prev_el {
			duplicate = append(duplicate, "_")
		}

		if el != prev_el {
			new_arr = append(new_arr, el)
		}

	}

	return new_arr, duplicate
}
