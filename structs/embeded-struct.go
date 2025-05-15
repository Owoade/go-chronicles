package main

import "fmt"

type DOB struct {
	year 	int
	month 	int
	date 	int
}

type Person struct {
	name 	string
	phone 	string
	dob		DOB
}

func main (){

	person := Person{
		name: "Owoade Anuoluwapo",
		phone: "0907XXXXXXX",
		dob: DOB{
			year: 0,
			month: 0,
			date: 0,
		},
	}

	fmt.Println(person)

}