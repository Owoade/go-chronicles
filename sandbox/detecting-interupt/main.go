package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	var ws sync.WaitGroup
	ws.Add(1)
	go func() {
		defer ws.Done()
		sig := <-c
		fmt.Println(sig, "interupt detected, closing program")
	}()
	ws.Wait()

}
