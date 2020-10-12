package main

import (
	"fmt"
)

func main() {
	var f []float32

	// Appends elements to the end of the slice
	// Append returns an updated slice. Therefore it is necessary to store the result
	// of append in a variable
	f = append(f, 1.2)
	f = append(f, 2.4, 4.8, 8.16)

	fmt.Println(f)

	// Appends elements of a slice, into another slice with the use of ...(ellipsis)
	f = append(f, f...)
	fmt.Println(f)
}
