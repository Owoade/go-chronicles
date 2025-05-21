package main

import (
	"errors"
	"fmt"
	"iter"
)

const (
	BadRequest = iota + 400
	UnAuthorized
	PaymentRequired
	Forbidden
	NotFound
	MethodNotAllowed
)

var errorMap = map[int]error{
	BadRequest:       errors.New("The server could not understand the request due to invalid syntax."),
	UnAuthorized:     errors.New("You are not authorized to access this resource. Please provide valid"),
	PaymentRequired:  errors.New("Payment is required to access this resource. Please complete the payment to proceed."),
	Forbidden:        errors.New("You do not have permission to access this resource."),
	NotFound:         errors.New("The requested resource could not be found. Please check the URL and try again."),
	MethodNotAllowed: errors.New("The HTTP method used is not allowed for this resource. Please check the allowed methods."),
}

func app(status int) error {

	if errorMap[status] != nil {
		return errorMap[status]
	} else {
		return nil
	}

}

func runtime() iter.Seq[error] {
	return func(yield func(error) bool) {
		states := []int{200, 400, 200, 301, 400, 401, 200, 405, 403, 200, 404}

		for _, state := range states {
			yield(app(state))
		}
	}
}

func main() {

	for err := range runtime() {
		if errors.Is(err, errorMap[BadRequest]) {
			fmt.Println("It's a bad request 🔴")
		} else if errors.Is(err, errorMap[UnAuthorized]) {
			fmt.Println("Where's your ID fam? 🔴")
		} else if errors.Is(err, errorMap[PaymentRequired]) {
			fmt.Println("Dude you gatta pay up man, Cash or card 💁🏼")
		} else if errors.Is(err, errorMap[Forbidden]) {
			fmt.Println("Some's not right 🧐")
		} else if errors.Is(err, errorMap[NotFound]) {
			fmt.Println("Not found chief")
		} else if errors.Is(err, errorMap[MethodNotAllowed]) {
			fmt.Println("Wrong method")
		} else {
			fmt.Println("We cool 🟢")
		}
	}

}
