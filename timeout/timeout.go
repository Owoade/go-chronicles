package main

import (
	"fmt"
	"time"
)

func timeOutRoutine(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Ori yo mi"
}

func main() {

	channel := make(chan string, 1)
	go timeOutRoutine(channel)

	select {
	case message := <-channel:
		fmt.Println("Routine Message: ", message)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout Message: Asiko ti koja")
	}
}
