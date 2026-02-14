package main

import (
	"fmt"
)

func sender(c chan string) {
	c <- "Please touch grass"
	fmt.Println("Sender: Moving on.....")
}

func main() {

	c := make(chan string)

	for range 6 {
		go sender(c)
	}
	counter := 0
	for message := range c {
		counter++
		fmt.Println(message)
		if counter == 2 {
			close(c)
		}
	}

}
