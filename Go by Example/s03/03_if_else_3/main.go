package main

import (
	"fmt"
)

func main() {
	serverStatusOK := true

	if serverStatusOK {
		fmt.Println("Server is up & running.")
	}

	if s := "FX"; serverStatusOK {
		fmt.Printf("%s server is up & running.", s)
	}

	// fmt.Println(s)	//compiler error: undefined
}
