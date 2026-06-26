package main

import "fmt"

func main() {
	person := "seyi"
	switch {
	case person == "seyi":
		fmt.Printf("%s is a very wierd guy", person)
	case person == "mayowa":
		fmt.Printf("%s is a very gentle guy", person)
	case person == "samuel":
		fmt.Printf("%s is in between", person)
	}
}
