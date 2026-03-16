package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RandomUser struct {
	Name struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

type RandomUserAPIResponse struct {
	Results []RandomUser `json:"results"`
}

func main() {
	uri := "https://randomuser.me/api/"

	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var apiResponse RandomUserAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		panic(err)
	}

	fmt.Println(apiResponse.Results[0])

}
