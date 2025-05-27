package main

import "fmt"

func listen[T comparable](channel <-chan T, callback func(arg T)) {
	event := <-channel
	callback(event)
}

func emit[T comparable](channel chan<- T, val T) {
	channel <- val
}

func main() {

	channel := make(chan bool, 1)

	emit(channel, true)

	listen(channel, func(status bool) {
		fmt.Println("Callback executed with value:", status)
	})
}
