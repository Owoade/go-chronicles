package main

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementCounter()
	go incrementCounter()
	wg.Wait()
	println("final counter: ", counter)
}

func incrementCounter() {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}
