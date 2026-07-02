package requestparser

import "strings"

type RequestInfo struct {
	Method      string
	Path        string
	HttpVersion string
}

func ParseRequestLine(line string) RequestInfo {
	var info RequestInfo

	fields := strings.SplitN(line, " ", 3)
	if len(fields) == 3 {
		info.Method = fields[0]
		info.Path = fields[1]
		info.HttpVersion = fields[2]
	}

	return info
}
