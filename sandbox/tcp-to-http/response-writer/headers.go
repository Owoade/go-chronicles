package response

import (
	"fmt"
	"strings"
)

type StatusCode int

const (
	// 1xx Informational
	StatusContinue           StatusCode = 100
	StatusSwitchingProtocols StatusCode = 101
	StatusProcessing         StatusCode = 102

	// 2xx Success
	StatusOK             StatusCode = 200
	StatusCreated        StatusCode = 201
	StatusAccepted       StatusCode = 202
	StatusNoContent      StatusCode = 204
	StatusPartialContent StatusCode = 206

	// 3xx Redirection
	StatusMovedPermanently  StatusCode = 301
	StatusFound             StatusCode = 302
	StatusSeeOther          StatusCode = 303
	StatusNotModified       StatusCode = 304
	StatusTemporaryRedirect StatusCode = 307
	StatusPermanentRedirect StatusCode = 308

	// 4xx Client Errors
	StatusBadRequest           StatusCode = 400
	StatusUnauthorized         StatusCode = 401
	StatusForbidden            StatusCode = 403
	StatusNotFound             StatusCode = 404
	StatusMethodNotAllowed     StatusCode = 405
	StatusConflict             StatusCode = 409
	StatusGone                 StatusCode = 410
	StatusUnsupportedMediaType StatusCode = 415
	StatusTooManyRequests      StatusCode = 429

	// 5xx Server Errors
	StatusInternalServerError StatusCode = 500
	StatusNotImplemented      StatusCode = 501
	StatusBadGateway          StatusCode = 502
	StatusServiceUnavailable  StatusCode = 503
	StatusGatewayTimeout      StatusCode = 504
)

var StatusText = map[StatusCode]string{
	// 1xx
	StatusContinue:           "Continue",
	StatusSwitchingProtocols: "Switching Protocols",
	StatusProcessing:         "Processing",

	// 2xx
	StatusOK:             "OK",
	StatusCreated:        "Created",
	StatusAccepted:       "Accepted",
	StatusNoContent:      "No Content",
	StatusPartialContent: "Partial Content",

	// 3xx
	StatusMovedPermanently:  "Moved Permanently",
	StatusFound:             "Found",
	StatusSeeOther:          "See Other",
	StatusNotModified:       "Not Modified",
	StatusTemporaryRedirect: "Temporary Redirect",
	StatusPermanentRedirect: "Permanent Redirect",

	// 4xx
	StatusBadRequest:           "Bad Request",
	StatusUnauthorized:         "Unauthorized",
	StatusForbidden:            "Forbidden",
	StatusNotFound:             "Not Found",
	StatusMethodNotAllowed:     "Method Not Allowed",
	StatusConflict:             "Conflict",
	StatusGone:                 "Gone",
	StatusUnsupportedMediaType: "Unsupported Media Type",
	StatusTooManyRequests:      "Too Many Requests",

	// 5xx
	StatusInternalServerError: "Internal Server Error",
	StatusNotImplemented:      "Not Implemented",
	StatusBadGateway:          "Bad Gateway",
	StatusServiceUnavailable:  "Service Unavailable",
	StatusGatewayTimeout:      "Gateway Timeout",
}

type HeaderPayload struct {
	Status        StatusCode
	ContentLength int
	ContentType   string
}

func GetHeader(p HeaderPayload) string {
	lines := []string{}
	lines = append(lines, fmt.Sprintf("HTTP/1.1 %d %s", p.Status, StatusText[p.Status]))
	lines = append(lines, fmt.Sprintf("Content-Length: %d", p.ContentLength))
	lines = append(lines, fmt.Sprintf("Content-Type: %s", p.ContentType))
	return strings.Join(lines, "\r\n")
}
