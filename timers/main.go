package main

import (
	"fmt"
	"time"
)

func cron(secs int, callback func()) *time.Timer {
	timer := time.NewTimer(time.Duration(secs) * time.Second)

	defer func() {
		<-timer.C
		callback()
	}()

	return timer
}

func main() {

	cron(2, func() {
		fmt.Println("This is a cron handler")
	})

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 clocked")
	}()

	time.Sleep(time.Second * 1)

	stop := timer2.Stop()

	if stop {
		fmt.Println("Timer 2 stopped prematurely")
	}

}
