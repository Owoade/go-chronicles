package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Print struct {
	statement string
}

func (p *Print) Run() {
	fmt.Println(p.statement)
}

var wg sync.WaitGroup

func main() {
	pool := NewPool(3)

	wg.Add(3)

	for range 3 {
		go func() {
			defer wg.Done()
			pool.AddTask(&Print{
				statement: fmt.Sprintf("random value: %d", rand.Int()),
			})
		}()
	}

	wg.Wait()
	pool.Shutdown()
}
