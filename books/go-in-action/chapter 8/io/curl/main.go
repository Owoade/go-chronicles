package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	fileName := os.Args[2]

	if url == "" {
		log.Panic("url not provided!")
	}

	if fileName == "" {
		log.Panic("file not provided")
	}

	res, err := http.Get(url)
	if err != nil {
		log.Panicf("http error: %v", err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Panicf("file creation error: %v", err)
	}

	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, res.Body)
	if err := res.Body.Close(); err != nil {
		log.Panicf("file closure error: %v", err)
	}
}
