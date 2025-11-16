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

}
