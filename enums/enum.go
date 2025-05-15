package main
import "fmt"

const (
	BadRequest = iota + 400
	UnAuthorized
	PaymentRequired
	Forbidden
	NotFound
	MethodNotAllowed
)

var errorMessage = map[int]string{
	BadRequest: "The server could not understand the request due to invalid syntax.",
	UnAuthorized: "You are not authorized to access this resource. Please provide valid",
	PaymentRequired: "Payment is required to access this resource. Please complete the payment to proceed.",
	Forbidden: "You do not have permission to access this resource.",
	NotFound: "The requested resource could not be found. Please check the URL and try again.",
	MethodNotAllowed: "The HTTP method used is not allowed for this resource. Please check the allowed methods.",
}

func main(){
	fmt.Println(errorMessage[400], errorMessage[UnAuthorized])
}