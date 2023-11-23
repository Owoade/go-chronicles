package main

import "fmt"

func main() {

	type User struct {
		username string
		password string
	}

	email_map := make(map[User]string)

	email_map[User{username: "makinde", password: "****"}] = "owoadeanu@checkouonce.com"

	email_map[User{username: "korede", password: "****"}] = "bayo.bayero@gmail.com"

	fmt.Println(email_map)

}
