package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func handleConnection(c *Conn) {
	defer c.Close()
	log.Printf("client connected: %d", c.RemoteAddr())

	reader := bufio.NewReader(c)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error reading request: %s", err)
			return
		}
		if err := handleRequest(message, c); err != nil {
			fmt.Fprintf(c, "request error: %s \n", err)
		}
	}
}

func handleRequest(req string, c *Conn) error {
	args := strings.Split(req, " ")
	if len(args) == 0 {
		return errors.New("no command provided")
	}

	command := args[0]
	command = strings.ToUpper(command)
	println("command: ", command)
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
		handleSetRequest(req, c)
	case "SETEX":
		if len(args) < 4 {
			return errors.New("key, value or expiry missing")
		}
		handleSetExRequest(req, c)
	case "BEGIN":
		handleBeginTx(c)
	case "ROLLBACK":
		handleRollbackTx(c)
	case "COMMIT":
		handleCommitTx(c)
	default:
		return errors.New("invalid command")
	}

	return nil
}

func handleGetRequest(req string, c *Conn) {
	args := strings.Split(req, " ")
	if len(args) == 0 {
		c.Write([]byte("no key provided"))
		return
	}

	if c.TransactionIsActive {
		v, ok := c.TransactionState.Store[strings.TrimSpace(args[0])]
		if ok {
			respond(c, v.raw)
			return
		}
	}

	mu.RLock()
	defer mu.RUnlock()
	v, ok := store[strings.TrimSpace(args[0])]
	if !ok {
		c.Write([]byte("nil \n"))
		return
	}

	respond(c, v.raw)
}

func handleSetRequest(req string, c *Conn) {
	rawRequest := req
	args := strings.Split(req, " ")[1:]
	key, value := strings.TrimSpace(args[0]), strings.TrimSpace(args[1])

	if c.TransactionIsActive {
		c.TransactionState.Commands = append(c.TransactionState.Commands, rawRequest)
		c.TransactionState.Store[key] = Value{
			raw: value,
		}
		respond(c, "SET")
		return
	}

	mu.Lock()
	defer mu.Unlock()
	store[key] = Value{
		raw: value,
	}
	respond(c, "SET")
}

func handleSetExRequest(req string, c *Conn) {
	rawRequest := req
	args := strings.Split(req, " ")[1:]

	key, value, expire := strings.TrimSpace(args[0]), strings.TrimSpace(args[1]), strings.TrimSpace(args[2])
	ttl, err := strconv.Atoi(expire)
	if err != nil {
		c.Write([]byte("invalid expiry value"))
		return
	}

	if c.TransactionIsActive {
		c.TransactionState.Commands = append(c.TransactionState.Commands, rawRequest)
		c.TransactionState.Store[key] = Value{
			raw: value,
		}
		respond(c, "SETEX")
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

	respond(c, "SETEX")
}

func handleBeginTx(c *Conn) {
	if c.TransactionIsActive {
		respond(c, "transaction started already")
		return
	}

	c.TransactionIsActive = true
	c.TransactionState = TransactionState{
		Store:    make(map[string]Value),
		Commands: make([]string, 0),
	}

	respond(c, "BEGIN")
}

func handleCommitTx(c *Conn) {
	if !c.TransactionIsActive {
		respond(c, "no transaction found")
		return
	}

	for _, cmd := range c.TransactionState.Commands {
		handleRequest(cmd, c)
	}

	c.TransactionIsActive = false
	c.TransactionState = TransactionState{}

	respond(c, "COMMIT")
}

func handleRollbackTx(c *Conn) {
	if !c.TransactionIsActive {
		respond(c, "no transaction found")
		return
	}

	c.TransactionIsActive = false
	c.TransactionState = TransactionState{}

	respond(c, "ROLLBACK")
}

func removeKey(key string) func() {
	return func() {
		mu.Lock()
		defer mu.Unlock()
		delete(store, key)
	}
}

func respond(c *Conn, message string) {
	fmt.Fprintf(c, "%s\n", message)
}
