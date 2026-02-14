package main

import (
	"fmt"
)

func search[T comparable](arr []T, el T) bool {

	for _, val := range arr {
		if val == el {
			return true
		}
	}

	return false
}

func main() {

	var articles = make(map[string][]string)

	var str string 



	if str == "" {
		fmt.Println(len(articles["BTC"]), str)
	}

	val := []int{10, 20, 30, 40, 50}
	println(val[0])
	newSlice := val[1:3]
	newSlice[0] = 2
	println(newSlice[0], val[1])

}
