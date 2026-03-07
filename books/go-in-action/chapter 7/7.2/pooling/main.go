package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(5)
	pool := NewPool(GetDbConnection, 2)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			performQuery(pool, "SELECT * FROM users;", i)
		}()
	}

	wg.Wait()
}

func performQuery(p *Pool, query string, i int) {
	fmt.Printf("%d running query: %s\n", i, query)
	time.Sleep(time.Duration(i) * time.Second)
	conn, err := p.Acquire()
	if err != nil {
		return
	}
	defer p.Release(conn)
}
