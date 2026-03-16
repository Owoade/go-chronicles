package main

import "log"

func init() {
	log.SetPrefix("GO-CHRONICLES-LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("This is a log")
}
