package main

import (
	"fmt"
	"time"

	"github.com/owoade/go-chronicles/package/helper"
	tinytime "github.com/wagslane/go-tinytime"
)

func main() {

	helper.SayHello()

	milliseconds := time.Now().UnixNano() / int64(time.Microsecond)

	time := tinytime.New(uint32(milliseconds))

	fmt.Println(time)

}
