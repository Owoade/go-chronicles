package main

import (
	"fmt"
	"time"
)

func main() {

	f("direct")

	go f("routine")

	f("direct")

	time.Sleep(time.Second / 8)

}

func f(method string) {
	fmt.Printf("Called %s\n", method)
}
