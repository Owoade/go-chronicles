package customparser

import (
	"bytes"
	"strings"
)

func ParseRequestHeader(lines [][]byte) map[string]any {
	headers := make(map[string]any)

	for _, line := range lines {
		parts := bytes.SplitN(line, []byte(":"), 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(string(parts[0]))
			value := strings.TrimSpace(string(parts[1]))
			headers[strings.ToLower(key)] = value
		}
	}

	return headers
}
