package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("My name is {{.}}\n")
	if err != nil {
		panic(err) // or handle error gracefully
	}

	t1 = template.Must(t1.Parse("My name is {{.}}\n")) // panics if .Parse() returns an error
	t1.Execute(os.Stdout, "Owoade")
	t1.Execute(os.Stdout, "Anuoluwapo")
	t1.Execute(os.Stdout, "Ayoade")

}
