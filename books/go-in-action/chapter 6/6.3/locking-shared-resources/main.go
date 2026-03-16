package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

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
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
