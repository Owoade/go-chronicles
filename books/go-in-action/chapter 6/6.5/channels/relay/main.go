package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4)
	race := make(chan int)
	go Runner(race, 1)
	race <- 0
	wg.Wait()
}

func Runner(c chan int, id int) {
	defer wg.Done()
	fmt.Printf("runner %d wating for baton \n", id)
	for {
		m, _ := <-c
		if m > 0 { fmt.Printf("runner %d received baton from %d\n", id, m) }
		if id < 4 {
			go Runner(c, id+1)
			c <- id
		}

		fmt.Printf("runner %d done \n", id)
		return
	}
}
