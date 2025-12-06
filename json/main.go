package main

import "encoding/json"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Wallet struct {
	Address string
	Balance int
}

func main() {
	person := Person{
		Name: "John",
		Age:  30,
	}

	wallet := Wallet{
		Address: "0x1234567890",
		Balance: 100,
	}

	personToJSON, _ := json.Marshal(person)
	walletToJSON, _ := json.Marshal(wallet)

	println(string(personToJSON))
	println(string(walletToJSON))

	personJSON := []byte(`{"name": "John", "age": 30}`)

	var personFromJSON Person
	json.Unmarshal(personJSON, &personFromJSON)

	var personMap map[string]any
	json.Unmarshal(personJSON, &personMap)

	println(personMap["name"].(string))
	println(personFromJSON.Name)
	println(personFromJSON.Age)

}
