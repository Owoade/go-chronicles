package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Value struct {
	raw          string
	shouldExpire bool
	timer        time.Timer
}

var store = make(map[string]Value)
var mu sync.RWMutex

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
		conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				slog.Info("shutting down server")
				return
			}
			log.Printf("error accepting connection: %s", err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(c net.Conn) {
	defer c.Close()
	log.Printf("client connected: %d", c.RemoteAddr())

	reader := bufio.NewReader(c)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error reading request: %s", err)
			continue
		}
		handleRequest(message, c)
	}
}

func handleRequest(req string, c net.Conn) error {
	args := strings.Split(req, " ")
	if len(args) == 0 {
		return errors.New("no command provided")
	}

	command := args[0]
	command = strings.ToUpper(command)
	switch command {
	case "GET":
		if len(args) < 2 {
			return errors.New("key missing")
		}
		handleGetRequest(strings.Join(strings.Split(req, " ")[1:], " "), c)
	case "SET":
		if len(args) < 3 {
			return errors.New("key or value missing")
		}
		handleSetRequest(strings.Join(strings.Split(req, " ")[1:], " "), c)
	case "SETEX":
		if len(args) < 4 {
			return errors.New("key, value or expiry missing")
		}
		handleSetExRequest(strings.Join(strings.Split(req, " ")[1:], " "), c)
	default:
		return errors.New("invalid command")
	}

	return nil
}

func handleGetRequest(req string, c net.Conn) {
	args := strings.Split(req, " ")
	if len(args) == 0 {
		c.Write([]byte("no key provided"))
		return
	}
	println(args[0])
	mu.RLock()
	defer mu.RUnlock()
	v, ok := store[args[0]]
	if !ok {
		c.Write([]byte("nil"))
		return
	}

	c.Write([]byte(v.raw))
}

func handleSetRequest(req string, c net.Conn) {
	args := strings.Split(req, " ")
	fmt.Println(args)
	if len(args) < 2 {
		c.Write([]byte("no key and value provided"))
		return
	}

	key, value := args[0], args[1]
	mu.Lock()
	defer mu.Unlock()
	store[key] = Value{
		raw: value,
	}
	c.Write([]byte("SET"))
}

func handleSetExRequest(req string, c net.Conn) {
	args := strings.Split(req, " ")
	if len(args) < 3 {
		c.Write([]byte("incomplete args"))
		return
	}

	key, value, expire := args[0], args[1], args[2]
	ttl, err := strconv.Atoi(expire)
	if err != nil {
		c.Write([]byte("invalid expiry value"))
		return
	}

	if v, ok := store[key]; ok {
		if v.shouldExpire {
			v.timer.Stop()
		}
	}

	timer := time.AfterFunc(time.Duration(ttl)*time.Millisecond, removeKey(key))
	mu.Lock()
	defer mu.Unlock()
	store[key] = Value{
		raw:          value,
		shouldExpire: true,
		timer:        *timer,
	}

	c.Write([]byte("SETEX"))
}

func removeKey(key string) func() {
	return func() {
		mu.Lock()
		defer mu.Unlock()
		delete(store, key)
	}
}
