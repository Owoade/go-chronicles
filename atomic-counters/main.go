package main

import (
	"sync"
	"sync/atomic"
	"fmt"
)

func main(){
	var wg sync.WaitGroup
	var ops atomic.Uint64

	// Spin up 50 go routines 
	for range 50 {
		wg.Add(1)
		go func(){
			for range 1000 {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Ops: ", ops.Load())
}

