package response

import (
	"encoding/json"
	"log"
	"net"
	"strings"
)

func WriteJSON(c net.Conn, StatusCode StatusCode, body string) {
	header := GetHeader(HeaderPayload{
		Status:        StatusCode,
		ContentLength: len(body),
		ContentType:   "application/json",
	})

	if err := json.Unmarshal([]byte(body), new(struct{})); err != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}

	lines := []string{header, body}
	c.Write([]byte(strings.Join(lines, "\r\n\r\n")))
}
