package main

import "fmt"

type notifier interface {
	notify()
}

type User struct {
	name string
	email string
}

func (u *User) notify(){
	fmt.Printf("sending notification to user email: %s, name: %s", u.name, u.email)
}

func sendNotification(n notifier){
	n.notify()
}

func main(){
	user := User{
		name: "owoade anu",
		email: "owoadeanu@yahoo.com",
	}

	sendNotification(&user)
}