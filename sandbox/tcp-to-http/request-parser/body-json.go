package requestparser

import "encoding/json"

func ParseJSONBody(b []byte) (map[string]any, error) {
	body := make(map[string]any)

	if err := json.Unmarshal(b, &body); err == nil {
		return body, err
	}

	return body, nil
}
