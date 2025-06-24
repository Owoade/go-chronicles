package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter map[string]int
	mutex   sync.Mutex
}

func (c *Counter) operation(key string, by int, wg *sync.WaitGroup) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counter[key] += by
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	const iterations = 10

	c := Counter{
		counter: map[string]int{
			"anxiety":   10,
			"badEnergy": 10,
			"hope":      0,
		},
	}

	for range iterations {
		wg.Add(3)
		go c.operation("hope", 200, &wg)
		go c.operation("badEnergy", -10, &wg)
		go c.operation("anxiety", -10, &wg)
	}

	wg.Wait()

	fmt.Println("Counter", c.counter)

}
