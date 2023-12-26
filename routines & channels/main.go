package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan int, 4)

	chan2 := make(chan int, 5)

	go routine_1(chan1)
	go routine_2(chan2)

	channel_value := <-chan1

	channel_value_2 := <-chan2

	fmt.Println(channel_value)

	fmt.Println((channel_value_2))

}

func routine_1(ch chan int) {

	ch <- int(time.Now().UnixNano() / int64(time.Millisecond))

}

func routine_2(ch chan int) {
	ch <- int(time.Now().UnixNano() / int64(time.Millisecond))
}
