package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan int) {
	for job := range jobs {
		fmt.Println("Job", job, "started by worker", id)
		time.Sleep(time.Second)
		fmt.Println("Job", job, "endded by worker", id)
		result <- job * 2
	}
}

func main() {

	jobs := make(chan int, 5)
	result := make(chan int)

	for i := 1; i <= 1; i++ {
		go worker(i, jobs, result)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 1; i <= 5; i++ {
		fmt.Println("Result from worker", <-result)
	}

}
