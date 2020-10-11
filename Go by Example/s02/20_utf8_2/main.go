package main

import (
	"fmt"
)

func main() {

	s := "こんにちは"

	for i, v := range s {
		fmt.Printf("%#U starts at byte position %d\n", v, i)
	}

	fmt.Println([]byte(s))

	fmt.Println()
	s = "𠜎ABCD"
	for i, v := range s {
		fmt.Printf("%#U starts at byte position %d\n", v, i)
	}
	fmt.Println([]byte(s))
}
