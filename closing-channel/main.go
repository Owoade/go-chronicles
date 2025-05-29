package main

import (
	"fmt"
	"time"
)

func worker(queue chan int) {
	for {
		job, more := <-queue

		if more {
			fmt.Println("Job recieved: ", job)
		} else {
			fmt.Println("No more jobs, worker going to sleep")
			return
		}
	}
}

func main() {

	queue := make(chan int)

	go worker(queue)

	for i := 1; i <= 5; i++ {
		queue <- i
	}

	close(queue)

	time.Sleep(time.Second)
}
