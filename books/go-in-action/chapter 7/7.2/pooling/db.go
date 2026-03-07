package main

import (
	"io"
)

type DbConnection struct {
	closed bool
}

func (conn *DbConnection) Close() error {
	conn.closed = true
	println("connection closed")
	return nil
}

func GetDbConnection() (io.Closer, error) {
	return &DbConnection{}, nil
}
