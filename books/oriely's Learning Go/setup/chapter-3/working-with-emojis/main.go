package main

import "fmt"

// substr is not set to “🌞  ” Instead, you get “�” That’s because you copied only the
// first byte of the sun emoji’s code point, which is not a valid code point on its own.
// slicing a string copies first byte of each char in a string, would be problematic for strings with chars having multiple bytes, like emojis
func main() {
	str := "🌞 The day is 🌞"
	substr := str[0:3]
	fmt.Println(substr, len(substr))
}
