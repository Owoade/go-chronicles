package main

import (
	"fmt"
	"slices"
)

var open = []rune{'{', '[', '('}
var close = []rune{')', ']', '}'}
var pHash = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

func ValidParenthensis(s string) bool {
	if len(s) < 2 {
		return false
	}
	var stack = []rune{}

	for _, b := range s {
		if slices.Contains(open, b) {
			stack = append(stack, b)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[0]
			if pHash[top] != b {
				return false
			}
			stack = stack[1:]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(ValidParenthensis("()[]{})"))
}
