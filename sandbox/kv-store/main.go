package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Value struct {
	raw          string
	shouldExpire bool
	timer        time.Timer
}

type TransactionState struct {
	Commands []string
	Store    map[string]Value
}

type Conn struct {
	net.Conn
	ID                  string
	TransactionIsActive bool
	TransactionState    TransactionState
}

var store = make(map[string]Value)
var mu sync.RWMutex

var (
	subscribers = make(map[string][]*Conn)
	subscribersMu sync.RWMutex
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	log.Println("KV store running!")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		listener.Close()
	}()

	for {
		_conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				slog.Info("shutting down server")
				return
			}
			log.Printf("error accepting connection: %s", err)
			continue
		}
		conn := NewConnection(_conn)
		go handleConnection(conn)
	}

}
