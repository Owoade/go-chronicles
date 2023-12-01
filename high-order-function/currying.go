package main

import "fmt"

func main() {
	product := mult(51)(6)

	fmt.Println(product)
}

func mult(num int) func(int) int {
	return func(sec_num int) int {
		return num * sec_num
	}
}
