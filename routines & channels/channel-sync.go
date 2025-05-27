package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan bool)
	channel2 := make(chan bool)

	go routine(channel)

	go routine2(channel2)

	fmt.Println("Main  routine started")

	if <-channel2 == false {
		fmt.Println("Second go routine said hello")
	}

	if true == <-channel {
		fmt.Println("Go routine said hi 🙋🏻")
	}

	fmt.Println("Main routine ended")

}

func routine(channel chan bool) {
	fmt.Println("Routine initiated")
	time.Sleep(time.Second)
	channel <- true
	fmt.Println("Routine complete")
}

func routine2(channel chan bool) {
	fmt.Println("Routine 2 initiated")
	channel <- false
	fmt.Println("Routine 2 complete")
}
