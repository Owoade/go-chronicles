package main

import "fmt"

func main() {
	name := "Anuoluwapo"

	if name == "Anuoluwapo" {
		fmt.Println("Anu is the best Developer in the world")
	} else {
		fmt.Println("🌚")
	}

	//  short hand
	if name := "Anuoluwapo"; name == "Anuoluwapo" {
		fmt.Println("Anu is the best Developer in the world")
	}
}
