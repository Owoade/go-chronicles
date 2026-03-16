package main

import (
	"encoding/json"
	"fmt"
)

type TobeMarhsaled struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var mapTobeMarshaled map[string]any = make(map[string]any)

func main() {
	// Populate map
	mapTobeMarshaled["Greatest Uite"] = "great!"
	mapTobeMarshaled["of the greatest"] = []string{"gba gba!", "shi won!"}

	structToBeMarshaled := TobeMarhsaled{
		Name: "Hello there",
		Age:  23,
	}

	structData, err := json.MarshalIndent(structToBeMarshaled, "", "	")
	if err != nil {
		panic(err)
	}

	mapData, err := json.MarshalIndent(mapTobeMarshaled, "", "	")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(structData), string(mapData))
}
