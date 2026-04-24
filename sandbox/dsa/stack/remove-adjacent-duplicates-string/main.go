package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RemoveAdjacentDuplicatesString("azxxzy"))
}

func RemoveAdjacentDuplicatesString(s string) string {
	var stack = []string{}
	for _, r := range s {
		if len(stack) > 0 {
			lastIn := stack[len(stack)-1]
			if lastIn == string(r) {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(r))
			}
		} else {
			stack = append(stack, string(r))
		}
	}

	return strings.Join(stack, "")
}
