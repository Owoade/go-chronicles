package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go worker("worker 1", c)
	go worker("worker 2", c)
	c <- "start"
	time.Sleep(10 * time.Second)
	time.Sleep(time.Minute / 2)
}

func worker(name string, c chan string) {
	for {
		_, ok := <-c
		if !ok {
			fmt.Printf("channel already closed, worker: %s \n", name)
			return
		}

		if rand.Intn(100)%7 == 0 {
			close(c)
			fmt.Printf("closing channel %s \n", name)
			return
		}
		fmt.Printf("worker running %s \n", name)
		time.Sleep(2 * time.Second)
		c <- "message"
	}
}
