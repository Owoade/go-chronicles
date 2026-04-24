package main

import (
	"fmt"
	"slices"
	"strings"
)

var alphaNum = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

func main() {
	fmt.Println(ValidPalindromeAlphaNumeric("ab@a"))
}

func ValidPalindromeAlphaNumeric(s string) bool {
	if len(s) < 2 {
		return true
	}

	var left = 0
	var right = len(s) - 1

	for left < right {
		leftChar := s[left]
		rightChar := s[right]

		if !slices.Contains(alphaNum, rune(leftChar)) {
			left++
			continue
		} else if !slices.Contains(alphaNum, rune(rightChar)) {
			right--
			continue
		} else {
			if strings.ToLower(string(leftChar)) != strings.ToLower(string(rightChar)) {
				return false
			}
			left++
			right--
		}
	}

	return true
}
