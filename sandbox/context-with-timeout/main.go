package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := make(chan string)
	go func() {
		c <- run(ctx)
	}()

	message := <-c
	println(message)

}

func run(ctx context.Context) string {
	i := 0
	for i < 5 {
		select {
		case <-ctx.Done():
			return "context canceled or timeout"
		default:
			log.Printf("Processor - Task #%d.", i)
			time.Sleep(time.Duration(i) * time.Second)
			i++
		}
	}
	return "run complete"
}
