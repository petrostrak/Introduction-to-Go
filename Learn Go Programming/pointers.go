package main

import (
	"fmt"
)

func main() {
	var a int = 31
	var b *int = &a
	fmt.Println(&a, *b)
}
