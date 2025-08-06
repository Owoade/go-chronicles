package main

import "fmt"

func sender(c chan string) {
	c <- "Please touch grass"
	fmt.Println("Sender: Moving on.....")
}

func main() {

	c := make(chan string, 5)

	for range 5 {
		go sender(c)
	}

	message := <-c

	fmt.Println(message)

}
