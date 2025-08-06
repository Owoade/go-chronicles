package main

import (
	"fmt"
	"strconv"
)

type Physique struct {
	height int
	weight int
}

type User struct {
	name        string
	age         int
	is_verified bool
	// nested struct
	Dob struct {
		day   int
		month int
		year  int
	}
	/*
			Embeded Struct
		 	If not passed as a pointer the struct will never be empty hence failing null check
	*/
	*Physique
}

type ServiceOpts struct {
	username string
	password string
	host     string
	port     int
}

// struct method
func (opts ServiceOpts) get_service_url() string {
	return opts.username + ":" + opts.password + "@" + opts.host + ":" + strconv.Itoa(opts.port)
}

func main() {

	user := User{
		name: "Owoade",
		age:  20,
	}

	user.Dob.day = 20

	user.Dob.month = 20

	user.Dob.year = 2000

	opts := ServiceOpts{
		host:     "https://checkout.once",
		port:     3020,
		username: "Owoade",
		password: "Dey play😂",
	}

	fmt.Println(opts.get_service_url())

	fmt.Println(user.name, user, user.Physique)

	anonymous := struct {
		id   int
		name string
	}{
		id:   1,
		name: "Owoade Anuoluwapo",
	}

	fmt.Println(anonymous)
}
