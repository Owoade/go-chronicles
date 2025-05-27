package main

import "fmt"

func main() {

	channel := make(chan string, 2)

	channel <- "buffered"

	channel <- "buffered_2"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

}
