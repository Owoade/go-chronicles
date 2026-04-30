package handler

import (
	"encoding/json"
	"net/http"
)

func Routes(){
	http.HandleFunc("/json", SendJSON)
}

func SendJSON(res http.ResponseWriter, req *http.Request) {
	jsonResponse := struct {
		Name  string
		Email string
	}{
		Email: "owoadeanu@yahoo.com",
		Name:  "Owoade Anu",
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(&jsonResponse)

}
