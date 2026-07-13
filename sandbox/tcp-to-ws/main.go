package main

import (
	"bufio"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"os/signal"
	"strings"
	"syscall"
)

const websocketGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	headers := map[string]string{}

	// Read request line
	_, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		fmt.Println(parts)

		headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	key := headers["Sec-WebSocket-Key"]
	if key == "" {
		return
	}

	accept := generateAcceptKey(key)

	response := fmt.Sprintf(
		"HTTP/1.1 101 Switching Protocols\r\n"+
			"Upgrade: websocket\r\n"+
			"Connection: Upgrade\r\n"+
			"Sec-WebSocket-Accept: %s\r\n"+
			"\r\n",
		accept,
	)

	conn.Write([]byte(response))

	for {
		msg, err := readFrame(reader)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Received:", string(msg))

		err = writeTextFrame(conn, string(msg))
		if err != nil {
			return
		}
	}
}

func generateAcceptKey(key string) string {
	hash := sha1.Sum([]byte(key + websocketGUID))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func readFrame(r *bufio.Reader) ([]byte, error) {

	var header [2]byte

	_, err := io.ReadFull(r, header[:])
	if err != nil {
		return nil, err
	}

	fin := header[0]&0x80 != 0
	opcode := header[0] & 0x0F

	if !fin {
		return nil, fmt.Errorf("fragmented frames not supported")
	}

	if opcode == 0x8 {
		return nil, io.EOF
	}

	masked := header[1]&0x80 != 0

	payloadLen := uint64(header[1] & 0x7F)

	switch payloadLen {
	case 126:
		var ext uint16
		if err := binary.Read(r, binary.BigEndian, &ext); err != nil {
			return nil, err
		}
		payloadLen = uint64(ext)

	case 127:
		if err := binary.Read(r, binary.BigEndian, &payloadLen); err != nil {
			return nil, err
		}
	}

	var maskingKey [4]byte

	if masked {
		_, err = io.ReadFull(r, maskingKey[:])
		if err != nil {
			return nil, err
		}
	}

	payload := make([]byte, payloadLen)

	_, err = io.ReadFull(r, payload)
	if err != nil {
		return nil, err
	}

	if masked {
		for i := range payload {
			payload[i] ^= maskingKey[i%4]
		}
	}

	return payload, nil
}

func writeTextFrame(w io.Writer, msg string) error {

	payload := []byte(msg)

	frame := []byte{
		0x81,
	}

	l := len(payload)

	switch {
	case l <= 125:
		frame = append(frame, byte(l))

	case l <= 65535:
		frame = append(frame, 126)
		tmp := make([]byte, 2)
		binary.BigEndian.PutUint16(tmp, uint16(l))
		frame = append(frame, tmp...)

	default:
		frame = append(frame, 127)
		tmp := make([]byte, 8)
		binary.BigEndian.PutUint64(tmp, uint64(l))
		frame = append(frame, tmp...)
	}

	frame = append(frame, payload...)

	_, err := w.Write(frame)
	return err
}
