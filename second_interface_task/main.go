package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file_name := os.Args[1] // this command reads file passed to the command go run
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file) // since 'File' type implements the 'Reader' interface, you will be able to reuse that io.Copy function
}
