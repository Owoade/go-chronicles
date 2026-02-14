package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func init(){
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./example2 <url>")
	}
}

func main(){
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, r.Body)
	if err = r.Body.Close(); err != nil {
		log.Fatal(err)
	}
}