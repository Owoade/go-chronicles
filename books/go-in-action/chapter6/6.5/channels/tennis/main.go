package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	court := make(chan int)
	wg.Add(2)
	go player("serena", court)
	go player("venus", court)
	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("%s won \n", name)
			return
		}

		n := rand.Intn(100)
		if n%73 == 0 {
			fmt.Printf("%s missed \n", name)
			close(court)
			return
		}

		fmt.Printf("%s hit \n", name)
		time.Sleep(2 * time.Second)
		court <- ball
	}
}
