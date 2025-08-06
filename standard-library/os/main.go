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
	errCopyFs := os.CopyFS("../os-copy", os.DirFS("."))
	if os.IsExist(errCopyFs) == true {
		fmt.Println("Path ../os-copy already exists")
	}
	os.RemoveAll("../os-copy")

	// Returns the path to the executable that started the program
	executablePath, _ := os.Executable()
	fmt.Println("Executable path", executablePath)

	// Environment variable
	fmt.Println(os.Getenv("HELLO"), os.Environ())
	value, _ := os.LookupEnv("owoade")
	fmt.Println(value)

	//creates a channel for writting and reading data from
	fileReader, fileWriter, _ := os.Pipe()

	go func() {
		fileWriter.Write([]byte("Hello there"))
	}()

	buffer2 := make([]byte, 3)
	fileReader.Read(buffer2)
	fmt.Println(buffer2)

	// Compares the metadata of two files
	file1, _ := os.Stat("test.txt")
	file2, _ := os.Stat("test2.txt")
	FILE_ARE_THE_SAME := os.SameFile(file1, file2)
	fmt.Println(FILE_ARE_THE_SAME)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
