package main

import (
	"fmt"
)

var symbols = []string{"[", "]"}
var numbers = []string{"1", "2", "2", "3", "4", "5", "6", "7", "8", "9"}

func DecodeString(s string) string {

}

func MultiplyString(s string, m int) string {
	val := ""
	for i := 0; i < m; i++ {
		val = val + s
	}
	return val
}

func main() {
	fmt.Println(DecodeString("2[abc]3[cd]ef"))
}
