package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(
			w,
			`{
				"status": true,
				"message": "Ok"
			}`,
		)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}
