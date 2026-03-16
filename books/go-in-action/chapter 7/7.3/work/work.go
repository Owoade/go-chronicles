package main

import (
	"fmt"
	"sync"
)

type Task interface {
	Run()
}

type Pool struct {
	closed         bool
	ch             chan Task
	maximumWorkers int
	wg             sync.WaitGroup
}

func NewPool(maximumWorkers int) *Pool {
	p := &Pool{
		ch: make(chan Task),
	}
	p.wg.Add(maximumWorkers)
	for i := 0; i < maximumWorkers; i++ {
		go func() {
			defer p.wg.Done()
			fmt.Printf("Worker %d active \n", i)
			for task := range p.ch {
				fmt.Printf("Worker %d running task \n", i)
				task.Run()
			}
		}()
	}

	return p
}

func (p *Pool) AddTask(t Task) {
	p.ch <- t
}

func (p *Pool) Shutdown() {
	p.wg.Wait()
	close(p.ch)
}
