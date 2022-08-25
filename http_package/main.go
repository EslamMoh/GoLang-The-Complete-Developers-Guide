package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// This is how we can create an empty slice and reserve specific space for it for the elemenets (99999)
	// We made this because we will use this slice with Read function and the read function cannot resize wth the amount of passed data
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs) // We call Body struct because it implements the reader, we know it has access to the read funciton
	// fmt.Println(string(bs))
	// Instead of the above, we will use the next line:
	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("This is a custom implementation for the writer interface")
	return len(bs), nil
}
