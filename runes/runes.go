package main

import (
    "fmt"
    "unicode/utf8"
)

func rangeFunc( str string ){
	for _, char := range str {
		// prints unicode value
		decodedRune := string(char)
		fmt.Println(char, decodedRune)
	}
}

func forLoop( str string ){
	for i:=0; i < len(str); i++ {
		// prints unicode value

		fmt.Println(str[i])
	}
}

func main(){

	str := "Owoade"

	runeCount := utf8.RuneCountInString(str)

	fmt.Println(">>>> Rune count <<<<", runeCount)

	fmt.Println(">>>> Range Iteration <<<<")

	rangeFunc(str)

	fmt.Println(">>>> For loop <<<<")

	forLoop(str)

}