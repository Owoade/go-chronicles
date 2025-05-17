package main

import "fmt"

func findIndex[E ~[]T, T comparable]( list E, el T ) int {
	
	for i:=0; i < len(list); i++ {
		
		item := list[i]

		if item == el {
			return i
		}
	}

	return -1
}

func main(){

	names := []string{"Anu", "Seyi", "Psalm"}

	seyiIndex := findIndex(names, "Seyi");

	psalmIndex := findIndex(names, "Psalm")

	fmt.Println(seyiIndex, psalmIndex)

}