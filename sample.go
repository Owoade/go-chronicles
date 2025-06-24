package main

import "fmt"

func  search[T comparable](arr []T, el T) bool {

	for _, val := range arr {
		if val == el {
			return true
		}
	}

	return false
}

func main(){

	stringElementIsFound := search([]string{"seyi", "is", "a", "seyi"}, "seyi")

	numberElementIsFound := search([]int{1,3,4,5}, 9)

	fmt.Println(stringElementIsFound, numberElementIsFound)

}