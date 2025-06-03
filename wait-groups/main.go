package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("Starting job", i)
	time.Sleep(time.Second)
	fmt.Println("Ended job", i)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	wg.Wait()
}
