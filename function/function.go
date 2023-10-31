package main

import (
	"fmt"
	"strings"
)

func sayMyname(name string) string {
	name = "Hey! nerd your name is\t" + name

	return name
}

func getNamesFromFullName(full_name string) (string, string) {

	names := strings.Split(full_name, " ")

	return names[0], names[1]

}

func main() {

	statement := sayMyname("Owoade")

	first_name, last_name := getNamesFromFullName("Owoade anuoluwapo")

	fmt.Println(first_name, last_name)

	fmt.Println(statement)

}
