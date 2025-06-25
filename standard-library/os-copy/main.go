package main

import (
	"fmt"
	"os"
)

func main() {
	// handling file operation
	file, err := os.Open("test.txt")
	check(err)
	buffer := make([]byte, 5)
	bytesRead, err := file.Read(buffer)
	check(err)
	fmt.Println(string(buffer[:bytesRead]))

	// changing Current working directory
	fmt.Println(os.Getwd())
	errorCh := os.Chdir("test")
	check(errorCh)
	fmt.Println(os.Getwd())
	os.Chdir("../")

	// File access and ownership
	errChmod := os.Chmod("test.txt", 0644)
	check(errChmod)

	errorChown := os.Chown("test.txt", 1, 1)
	check(errorChown)

	// Copies content of the cwd into the test dir
	// info, errDir := os.Stat("../os-copy")

	// if errDir != nil || !info.IsDir() {
	// 	os.Mkdir("../os-copy", 0644)
	// }

	os.CopyFS("../os-copy", os.DirFS("."))

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
