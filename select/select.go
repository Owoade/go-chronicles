package main

import (
	"fmt"
	"time"
)

func selectRoutine( channel chan<- string ){
	time.Sleep(2 * time.Second)
	channel <- "Hello there ✋🏿 from Salem"
}

func selectRoutine2( channel chan<- string ){
	time.Sleep(1 * time.Second)
	channel <- "Hello there ✋🏿 from the cave of adulam"
}

func main(){

	channel := make(chan string, 1);

	channel2 := make(chan string, 1);

	go selectRoutine(channel)

	go selectRoutine2(channel2)

	select {
		case message := <-channel:
			fmt.Println("message: ", message)
		case message2 := <-channel2:
			fmt.Println("message2: ", message2)
	}
}