package requestparser

import (
	"bytes"
	"strings"
)

type FormData struct {
	Key   string
	Type  string
	Bytes []byte
}

func ParseFormData(body []byte, boundary string) []FormData {
	var allFormData []FormData

	parts := bytes.Split(body, []byte(boundary))

	for _, part := range parts {
		part = bytes.Trim(part, "\r\n-")

		if len(part) == 0 {
			continue
		}

		// split headers and body ONCE
		split := bytes.SplitN(part, []byte("\r\n\r\n"), 2)
		if len(split) < 2 {
			continue
		}

		headerPart := split[0]
		bodyPart := split[1]

		var formData FormData

		headers := bytes.Split(headerPart, []byte("\r\n"))

		for _, h := range headers {
			line := string(h)

			if strings.HasPrefix(line, "Content-Disposition") {
				formData.Key = getKey(line)
			}

			if strings.HasPrefix(line, "Content-Type") {
				formData.Type = getContentType(line)
			}
		}

		if formData.Type == "" {
			formData.Type = "text"
		}

		// IMPORTANT: raw bytes only
		formData.Bytes = bodyPart

		allFormData = append(allFormData, formData)
	}

	return allFormData
}

func getKey(line string) string {
	parts := strings.Split(line, `Content-Disposition: form-data; name="`)
	if len(parts) > 1 {
		return strings.Split(parts[1], `"`)[0]
	}
	return ""
}

func getContentType(line string) string {
	parts := strings.Split(line, "Content-Type: ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
