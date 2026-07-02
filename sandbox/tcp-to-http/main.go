package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	requestparser "sandbox.tcptohttp/request-parser"
	"sandbox.tcptohttp/response-writer"
)

func main() {
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
		if string(headerSection) == "" {
			continue
		}

		var bodySection []byte
		if len(parts) == 2 {
			bodySection = parts[1]
		}

		lines := bytes.Split(headerSection, []byte("\r\n"))
		requestLine := string(lines[0])
		headerLines := lines[1:]
		_ = requestparser.ParseRequestLine(requestLine)
		headers := requestparser.ParseRequestHeader(headerLines)

		contentLength, _ := strconv.Atoi(headers["content-length"].(string))
		remaining := contentLength - len(bodySection)

	readMoreLoop:
		for remaining > 0 {
			n, err := conn.Read(tmp)

			if err != nil {
				break readMoreLoop
			}

			bodySection = append(bodySection, tmp[:n]...)
			remaining -= n
		}

		contentType := (headers["content-type"]).(string)
		println(contentType)

		if contentType == "application/json" {
			body, _ := requestparser.ParseJSONBody(bodySection)
			resoponseBody, _ := json.Marshal(body)
			response.WriteJSON(conn, response.StatusOK, string(resoponseBody))
		} else if strings.Contains(contentType, "multipart/form-data;") {
			contentTypeParts := strings.SplitN(contentType, "=", 2)
			boundary := contentTypeParts[1]
			parts := requestparser.ParseFormData(bodySection, boundary)
			for _, part := range parts {
				if part.Type == "text" {
					println(string(part.Bytes))
				} else {
					extension := mimeToExtension[part.Type]
					fileName := strconv.Itoa(int(time.Now().UnixMicro()))
					os.WriteFile(fmt.Sprintf("tmp/%s.%s", fileName, extension), part.Bytes, 0644)
				}
			}
		} else {
			fileName := strconv.Itoa(int(time.Now().UnixMicro()))
			extension := mimeToExtension[contentType]
			os.WriteFile(fmt.Sprintf("tmp/%s.%s", fileName, extension), bodySection, 0644)
		}

		conn.Close()

	}
}
