package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os/signal"
	"syscall"

	customparser "sandbox.tcptohttp/parser"
)

func main() {
	println("hello there \n")
	println("hello there \n")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	log.Println("tcp server running on port 8080")

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

		var buf []byte
		tmp := make([]byte, 4096)

	connLoop:
		for {
			n, err := conn.Read(tmp)
			buf = append(buf, tmp[:n]...)

			if n == 0 {
				break
			}

			if bytes.Contains(buf, []byte("\r\n\r\n")) {
				//end of header
				break connLoop
			}

			if err != nil {
				break connLoop
			}
		}

		parts := bytes.SplitN(buf, []byte("\r\n\r\n"), 2)
		headerSection := parts[0]
		var bodySection []byte
		if len(parts) == 2 {
			bodySection = parts[1]
		}

		lines := bytes.Split(headerSection, []byte("\r\n"))
		requestLine := string(lines[0])
		headerLines := lines[1:]
		requestInfo := customparser.ParseRequestLine(requestLine)
		headers := customparser.ParseRequestHeader(headerLines)

		contentType := headers["content-type"]

		if contentType == "application/json" {
			body, _ := customparser.ParseJSONBody(bodySection)
			fmt.Println(body)
		}

		fmt.Println(requestInfo, headers, string(bodySection))

		conn.Close()

	}
}
