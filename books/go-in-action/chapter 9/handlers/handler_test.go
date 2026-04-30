package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handler "go.in.action.9/handlers"
)

func init() {
	handler.Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Should test SendJSON func")
	req, err := http.NewRequest("GET", "/json", nil)
	if err != nil {
		t.Fatalf("error creating request: %s", err)
	}

	rec := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expecting %d, got %d", http.StatusOK, rec.Code)
	}

	t.Log("test passed")
}
