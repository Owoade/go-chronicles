package main

type User struct {
	name string
	email string
	phone string
	day int
	DOB
}

type DOB struct {
	day int
	month int
	year int
}

func (d DOB) Print(){

}

func main(){

	user := User{
		name: "owoade anu",
		email: "owoade@email.com",
		phone: "0901122345",
		day: 30,
		DOB: DOB{
			day: 10,
			month: 10,
			year: 2024,
		},
	}

	println(user.day, user.DOB.day)



	


}