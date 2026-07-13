package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// ============================================================================
// LEVEL 1: Simple Length-Prefix Framing (no bit manipulation)
// ============================================================================

// Frame format: [4-byte length][payload]
// Example: "Hello" -> [0,0,0,5]['H','e','l','l','o']

func writeSimpleFrame(w io.Writer, msg []byte) error {
	// Write 4-byte length prefix (big-endian)
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(msg)))

	_, err := w.Write(length)
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	return err
}

func readSimpleFrame(r *bufio.Reader) ([]byte, error) {
	// Read 4-byte length
	lengthBytes := make([]byte, 4)
	_, err := io.ReadFull(r, lengthBytes)
	if err != nil {
		return nil, err
	}

	// Convert to uint32
	length := binary.BigEndian.Uint32(lengthBytes)

	// Read exactly 'length' bytes
	payload := make([]byte, length)
	_, err = io.ReadFull(r, payload)
	return payload, err
}

// ============================================================================
// LEVEL 2: Add a Message Type (like WebSocket opcodes)
// ============================================================================

// Frame format: [1-byte type][4-byte length][payload]
// Example: Text message "Hi" -> [1,0,0,0,2,'H','i']

type MessageType byte

const (
	TypeText MessageType = 1
	TypeBinary MessageType = 2
	TypeClose MessageType = 3
)

func writeTypedFrame(w io.Writer, msgType MessageType, msg []byte) error {
	// Write type byte
	if _, err := w.Write([]byte{byte(msgType)}); err != nil {
		return err
	}

	// Write length prefix
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(msg)))
	if _, err := w.Write(length); err != nil {
		return err
	}

	// Write payload
	_, err := w.Write(msg)
	return err
}

func readTypedFrame(r *bufio.Reader) (MessageType, []byte, error) {
	// Read type byte
	typeByte := make([]byte, 1)
	_, err := io.ReadFull(r, typeByte)
	if err != nil {
		return 0, nil, err
	}

	// Read length
	lengthBytes := make([]byte, 4)
	_, err = io.ReadFull(r, lengthBytes)
	if err != nil {
		return 0, nil, err
	}
	length := binary.BigEndian.Uint32(lengthBytes)

	// Read payload
	payload := make([]byte, length)
	_, err = io.ReadFull(r, payload)
	return MessageType(typeByte[0]), payload, err
}

// ============================================================================
// LEVEL 3: Add XOR Masking (like WebSocket)
// ============================================================================

// Frame format: [1-byte type][1-byte mask?][4-byte length][4-byte mask][payload]

func writeMaskedFrame(w io.Writer, msgType MessageType, msg []byte, maskKey []byte) error {
	// Write type byte with mask flag
	maskFlag := byte(0)
	if len(maskKey) == 4 {
		maskFlag = 0x80
	}
	if _, err := w.Write([]byte{byte(msgType) | maskFlag}); err != nil {
		return err
	}

	// Write length
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(msg)))
	if _, err := w.Write(length); err != nil {
		return err
	}

	// Write mask if present
	if len(maskKey) == 4 {
		if _, err := w.Write(maskKey); err != nil {
			return err
		}

		// Mask the payload
		masked := make([]byte, len(msg))
		for i := range msg {
			masked[i] = msg[i] ^ maskKey[i%4]
		}
		_, err := w.Write(masked)
		return err
	}

	_, err := w.Write(msg)
	return err
}

func readMaskedFrame(r *bufio.Reader) (MessageType, []byte, error) {
	// Read type byte
	typeByte := make([]byte, 1)
	_, err := io.ReadFull(r, typeByte)
	if err != nil {
		return 0, nil, err
	}

	msgType := MessageType(typeByte[0] & 0x7F) // Clear mask flag
	masked := (typeByte[0] & 0x80) != 0        // Check mask flag

	// Read length
	lengthBytes := make([]byte, 4)
	_, err = io.ReadFull(r, lengthBytes)
	if err != nil {
		return 0, nil, err
	}
	length := binary.BigEndian.Uint32(lengthBytes)

	// Read mask if present
	var maskKey []byte
	if masked {
		maskKey = make([]byte, 4)
		_, err = io.ReadFull(r, maskKey)
		if err != nil {
			return 0, nil, err
		}
	}

	// Read payload
	payload := make([]byte, length)
	_, err = io.ReadFull(r, payload)
	if err != nil {
		return 0, nil, err
	}

	// Unmask if needed
	if masked {
		for i := range payload {
			payload[i] ^= maskKey[i%4]
		}
	}

	return msgType, payload, nil
}

// ============================================================================
// Demo server using Level 2 (typed frames)
// ============================================================================

func handleEchoConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msgType, data, err := readTypedFrame(reader)
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}

		fmt.Printf("Received [%s]: %s\n", msgType, string(data))

		// Echo back
		if msgType == TypeText {
			writeTypedFrame(conn, TypeText, data)
		} else if msgType == TypeClose {
			writeTypedFrame(conn, TypeClose, nil)
			return
		}
	}
}

func main() {
	// Demo the framing functions
	fmt.Println("=== Framing Demo ===\n")

	// LEVEL 1: Simple framing
	fmt.Println("LEVEL 1: Simple length-prefix framing")
	msg := []byte("Hello, WebSocket!")
	fmt.Printf("Original: %s\n", msg)

	// Simulate round-trip
	buf := bufio.NewReader(nil) // placeholder

	_ = buf // avoid unused import error in demo

	// LEVEL 2: Typed framing
	fmt.Println("\nLEVEL 2: Typed frames")
	fmt.Printf("Text message type: %d\n", TypeText)
	fmt.Printf("Binary message type: %d\n", TypeBinary)
	fmt.Printf("Close message type: %d\n", TypeClose)

	// LEVEL 3: Masked framing
	fmt.Println("\nLEVEL 3: XOR masking")
	original := []byte("secret")
	mask := []byte{0x37, 0x84, 0x92, 0x15}
	masked := make([]byte, len(original))
	for i := range original {
		masked[i] = original[i] ^ mask[i%4]
	}
	fmt.Printf("Original: %v (%s)\n", original, original)
	fmt.Printf("Masked:   %v\n", masked)

	// Unmask to verify
	unmasked := make([]byte, len(masked))
	for i := range masked {
		unmasked[i] = masked[i] ^ mask[i%4]
	}
	fmt.Printf("Unmasked: %v (%s)\n", unmasked, unmasked)

	fmt.Println("\n=== To run the server: uncomment the server code ===")

	// Uncomment to run server:
	/*
		listener, err := net.Listen("tcp", ":9000")
		if err != nil {
			panic(err)
		}
		defer listener.Close()
		fmt.Println("Server on :9000")

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go handleEchoConnection(conn)
		}
	*/
}