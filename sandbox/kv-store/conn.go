package main

import "net"

func NewConnection(c net.Conn) *Conn {
	return &Conn{
		Conn: c,
	}
}
