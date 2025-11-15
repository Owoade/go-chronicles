package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				write.resp <- true
			}
		}
	}()

	for range 100 {
		go func() {
			read := readOp{
				key:  rand.Intn(5),
				resp: make(chan int),
			}
			reads <- read
			<-read.resp
			atomic.AddUint64(&readOps, 1)
		}()
	}

	for range 50 {
		go func() {
			write := writeOp{
				key:   rand.Intn(5),
				value: rand.Intn(200),
				resp:  make(chan bool),
			}
			writes <- write
			<-write.resp
			atomic.AddUint64(&writeOps, 1)
		}()
	}

	time.Sleep(time.Second)

	totalReadOps := atomic.LoadUint64(&readOps)
	fmt.Println("Total read ops", totalReadOps)
	totalWriteOps := atomic.LoadUint64(&writeOps)
	fmt.Println("Total write ops", totalWriteOps)
}

