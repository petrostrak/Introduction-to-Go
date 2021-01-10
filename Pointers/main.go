package main

import (
	"fmt"
)

func main() {
	a := 25
	b := &a                   // translates to ADDRESS OF
	fmt.Printf("b:%T", b)     // *int
	fmt.Printf("\n&b:%T", &b) // *(*int)
	fmt.Printf("\n*b:%T", *b) // derefferencing b, so type int and value 25
}
