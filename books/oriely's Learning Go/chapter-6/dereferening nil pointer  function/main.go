package main

import "fmt"

func main() {
	var f *int = new(int)
	Update(f)
	fmt.Println(*f) // 20
}

func Update(v *int) {
	var x = 20
	*v = x
}
