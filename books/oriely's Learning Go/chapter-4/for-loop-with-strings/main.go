package main

import "fmt"

func main() {
	str := "Hello 👌 World 🌞"

	fmt.Println("loop: for range")
	for j, v := range str {
		// Processes each char appropriately
		fmt.Println(j, string(v))
	}

	fmt.Println("lopp: traditional loop")
	for i := 0; i < len(str); i++ {
		// Breaks each char into bytes, would provide unexpected behaviour for multi-bytes char
		fmt.Println(i, string(str[i]))
	}
}
