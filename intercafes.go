package main

import (
	"fmt"
)

func main() {
	var w Write = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Winter interface {
	Write([]byte) (int, error) // This method eccepts a slice of bytes and return an integer and a error
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
