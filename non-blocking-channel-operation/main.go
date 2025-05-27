package main

import "fmt"

func main() {

	channel := make(chan string, 1)

	channel2 := make(chan string)

	select {
	case channel <- "Anu said hello":
		fmt.Println("Message sent")
	case message := <-channel2:
		fmt.Println("Message recieved from second channel", message)
	default:
		fmt.Println("No activities")
	}

	select {
	case message := <-channel:
		fmt.Println("Message recieved from first channel", message)
	default:
		fmt.Println("No activities")
	}

}
