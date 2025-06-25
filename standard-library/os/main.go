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
	// Creates folder in case the folder doesn't exist
	os.CopyFS("../os-copy", os.DirFS("."))

	// Returns the path to the executable that started the program
	executablePath, _ := os.Executable()
	fmt.Println("Executable path", executablePath)

	fmt.Println(os.Getenv("HELLO"), os.Environ())

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
