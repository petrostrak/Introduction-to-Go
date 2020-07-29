package main

import (
	"fmt"
)

func main() {
	switch 9 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 9:
		fmt.Println("two, four or nine")
	default:
		fmt.Println("another number")
	}
}
