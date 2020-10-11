package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello𠜎ち"      // Hello(5)𠜎(4)ち(3) = 12
	fmt.Println(len(s)) // len() returns the number of bytes in a string
	fmt.Println(utf8.RuneCountInString(s))
}
