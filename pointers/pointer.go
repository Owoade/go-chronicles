package main

import "fmt"

type User struct {
	name         string
	age          int
	remark       string
	is_unserious bool
}

func (u *User) get_remark() {

	u.remark = "This is user is quite unserious, do better!"

}

func main() {

	name := "Makinde Mayowa"

	name_pointer := &name

	*name_pointer = "Mayor Mankind"

	fmt.Println(name, name_pointer)

	user := User{
		name:         "Olaoye Seyi",
		age:          12,
		remark:       "",
		is_unserious: true,
	}

	user.get_remark()

	fmt.Println((user))

}
