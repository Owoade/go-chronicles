package main

import "fmt"

// Shadowing
func main() {
	x := 10
	fmt.Println(x)
	{
		x := 20
		fmt.Println(x)
	}
	fmt.Println(x)

	// Shadowing imported packges (very bad)
	fmt := 300
	fmt.Println(x) // won't compile

	// Shadowing predeclared identifiers (very bad, can introduce silent bugs)
	true := 300
	println(true)
}
