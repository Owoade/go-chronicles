package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleSendJSON() {
	req, _ := http.NewRequest("GET", "/json", nil)

	rec := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rec, req)

	var body struct {
		Name  string
		Email string
	}
	if err := json.NewEncoder(rec).Encode(&body); err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
	// Output:
	// {owoadeanu@yahoo.com Owoade Anu}
}
