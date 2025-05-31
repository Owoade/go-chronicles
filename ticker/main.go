package main

import (
	"fmt"
	"time"
)

type IntervalReturn struct {
	value *time.Ticker
	stop func()
}

func interval( secs int, callback func() )(IntervalReturn){

	ticker := time.NewTicker(time.Duration(secs) * time.Second);

	done := make( chan bool );

	go func ()  {
		for {
			select {
			case <-ticker.C:
				callback()
			case <-done:
				return
			}
		}
	}()

	return IntervalReturn{
		value: ticker,
		stop: func() {
			done <- true
			ticker.Stop()
		},
	}

}

func main(){

	ticker := interval(1, func() {
		fmt.Println("Olaoye Seyi is a Senior dev")
	})

	time.Sleep(5*time.Second)

	ticker.stop()

	fmt.Println("Ticker stopped")
	
}