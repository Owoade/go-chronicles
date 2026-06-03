package main

import (
	"net"
	"strconv"
	"time"
)

func NewConnection(c net.Conn) *Conn {
	return &Conn{
		ID:   strconv.Itoa(int(time.Now().UnixMilli())),
		Conn: c,
	}
}
