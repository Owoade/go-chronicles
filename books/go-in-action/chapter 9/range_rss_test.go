package main

import (
	"net/http"
	"testing"
)

type TestCase struct {
	URL                string
	ExpectedStatusCode int
}

var testCases = []TestCase{
	{
		URL:                "https://smoothballot.com",
		ExpectedStatusCode: 200,
	},
	{
		URL:                "https://facebook.com",
		ExpectedStatusCode: 200,
	},
	{
		URL:                "https://api.smoothballot.com/random-route",
		ExpectedStatusCode: 404,
	},
}

func TestRangeDownload(t *testing.T) {
	for _, testCase := range testCases {
		t.Log("Giveing the need to download content")
		t.Logf("Checking %s for status code %d", testCase.URL, testCase.ExpectedStatusCode)

		resp, err := http.Get(testCase.URL)
		if err != nil {
			t.Fatalf("Unable to make get call to %s", testCase.URL)
		}
		t.Log("HTTP GET call successful")
		defer resp.Body.Close()

		if resp.StatusCode != testCase.ExpectedStatusCode {
			t.Fatalf("Expected status code %d, got %d", testCase.ExpectedStatusCode, resp.StatusCode)
		}

		t.Log("Test successful")
	}
}
