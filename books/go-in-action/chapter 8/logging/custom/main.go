package main

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	Trace = log.New(
		io.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Llongfile,
	)

	Info = log.New(
		os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Llongfile,
	)

	Warning = log.New(
		os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Llongfile,
	)

	Error = log.New(
		io.MultiWriter(
			os.Stderr,
			file,
		),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Llongfile,
	)
}

func main() {
	Trace.Println("Hello there!")
	Info.Println("Hello there!")
	Warning.Println("Hello there!")
	Error.Println("Hello there!")
}
