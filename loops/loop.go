package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 8}

	fmt.Println(get_max_element(arr))

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
