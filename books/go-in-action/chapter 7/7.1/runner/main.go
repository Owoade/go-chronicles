package main

import (
	"log"
	"time"
)

// Runner execute sets of task within a specified duration
// Should timeout once duration passes
// Interupt from os (Ctrl + S) should stop the program immediately

func main() {
	runner := NewRunner(5 * time.Second)
	runner.Add(
		createTask(),
		createTask(),
		createTask(),
		createTask(),
		createTask(),
	)
	if err := runner.Start(); err != nil {
		panic(err)
	}
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
