package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	requests := make(chan int, 5)

	for range 5 {
		requests <- rand.Int()
	}

	close(requests)

	for request := range requests {
		<-time.Tick(500 * time.Millisecond)
		fmt.Println("Request", request)
	}

	burstyLimiter := make(chan time.Time, 5)

	burstyRequests := make(chan int, 5)

	for range 5 {
		burstyRequests <- rand.Int()
	}

	close(burstyRequests)

	go func() {
		for value := range time.Tick(500 * time.Millisecond) {
			burstyLimiter <- value
		}
	}()

	for range 3 {
		burstyLimiter <- time.Now()
	}

	for burstyRequest := range burstyRequests {
		<-burstyLimiter
		fmt.Println("BurstyRequest", burstyRequest)
	}

}
