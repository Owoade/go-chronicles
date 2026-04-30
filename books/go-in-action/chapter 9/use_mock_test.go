package main

import (
	"net/http"
	"testing"
)

func TestMockServer(t *testing.T) {
	server := mockServer()
	defer server.Close()

	t.Log("Using mock server")
	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error calling url %s", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d got %d", res.StatusCode, http.StatusOK)
	}

	t.Log("Test passed")
}
