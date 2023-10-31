package main

import (
	"errors"
	"fmt"
)

func main() {

	str, err := error_function(4)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println((str))

}

// it returns an error if the parameter parsed is an odd number
func error_function(num int) (str string, err error) {

	if num%2 != 0 {
		err := errors.New("The value passed in is an odd number")
		return "", err
	}

	return "you passed in an even number🎉", nil

}
