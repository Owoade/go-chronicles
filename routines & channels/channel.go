package main

import "fmt"

func main() {

	channel := make(chan string)

	go test_finc("go routine", &channel)

	// This blocks execution
	message := <-channel

	fmt.Println(message)

	fmt.Println("This should be called last")

}

func test_finc(method string, channel *chan string) {
	*channel <- "Bawoni"
	fmt.Println("Called: ", method)
}
