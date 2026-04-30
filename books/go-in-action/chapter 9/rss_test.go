package main

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Giveing the need to download content")
	t.Logf("Checking %s for status code %d", url, statusCode)

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Unable to make get call to %s", url)
	}
	t.Log("HTTP GET call successful")
	defer resp.Body.Close()

	if resp.StatusCode != statusCode {
		t.Fatalf("Expected status code %d, got %d", statusCode, resp.StatusCode)
	}

	t.Log("Test successful")

}
