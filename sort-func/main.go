package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Person struct {
	name string
	age  int
}

func main() {

	persons := []Person{
		Person{
			name: "Seyi",
			age:  30,
		},
		Person{
			name: "Mayowa",
			age:  50,
		},
		Person{
			name: "Bidemi",
			age:  21,
		},
	}

	slices.SortFunc(persons, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println(persons)

}
