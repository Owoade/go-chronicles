package main

import (
	"fmt"
	"time"
)

func iterator( n int ) chan string {

	channel := make( chan string, n );

	for i := 1; i <= n; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Printf("Iteration after %ds \n", i)
		channel <- "Hello"
	} 

	// Due to the invovation of close() with the channel
	// the range iteration in the main function terminates after receiving the 2 elements.
	close(channel)

	return channel
	
}

func main() {

	for value := range iterator(5) {
		fmt.Println(value)
	}

}