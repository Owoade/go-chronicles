package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("Anu"))

	fmt.Fprintf(&b, " Owoade Ayoade")

	b.WriteTo(os.Stdout)

	copy()
}
