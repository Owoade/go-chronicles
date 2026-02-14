package main

import (
	"fmt"
	"strconv"
)

type Entity interface {
	Compute() string
}

type AnotherEntity struct{}

type GenZ struct {
	name          string
	year_of_birth int
}

type Millenial struct {
	name          string
	year_of_birth int
}

func (e AnotherEntity) Compute() string {
	return "It is well"
}

func (genz GenZ) Compute() string {
	if genz.year_of_birth < 1997 {
		return "No dey famz, you no be genZ"
	}

	genZ_age := genz.year_of_birth - 1997

	if genZ_age == 0 {
		return "You are a genZ baby"
	}

	return "You are " + strconv.Itoa(genZ_age) + " genZ years old"

}

func (millenial Millenial) Compute() string {

	if millenial.year_of_birth > 1996 {
		return "You no be millenial, go meet your mate"
	}

	millenial_age := 1997 - millenial.year_of_birth

	if millenial_age == 1 {
		return "Na God save you sha"
	}

	return "You are " + strconv.Itoa(millenial_age) + " years older than the oldest genZ"

}

func compute_entity(c Entity) string {
	return c.Compute()
}

func main() {

	var ae AnotherEntity

	plato := new(GenZ)
	plato.name = "Name"
	plato.year_of_birth = 2000

	fmt.Println(plato.Compute())

	fmt.Println(ae.Compute())

	fmt.Println(compute_entity(ae))

	fmt.Println(compute_entity(GenZ{
		name:          "Owoade",
		year_of_birth: 2004,
	}))

	println(compute_entity(&ae))

	fmt.Println(compute_entity(Millenial{
		name:          "Bonnie",
		year_of_birth: 1960,
	}))
	
}
