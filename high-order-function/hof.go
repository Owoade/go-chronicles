package main

import "fmt"

func main() {

	test_arr := []int{1, 2, 4, 5}

	even_vals := filter_num(test_arr, func(num int) bool {
		return num%2 == 0
	})

	fmt.Println(even_vals)

}

func filter_num(array_num []int, filter_callback func(num int) bool) []int {

	var new_arr []int

	for _, val := range array_num {

		is_valid := filter_callback(val)

		if is_valid {
			new_arr = append(new_arr, val)
		}
	}

	return new_arr

}
