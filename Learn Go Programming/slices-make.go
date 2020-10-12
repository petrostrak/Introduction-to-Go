package main

import (
	"fmt"
)

func main() {
	// a := make([]int, 3)
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length : %v\n", len(a))
	fmt.Printf("Capacity : %v\n", cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length : %v\n", len(a))
	fmt.Printf("Capacity : %v\n", cap(a))
	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Printf("Length : %v\n", len(a))
	fmt.Printf("Capacity : %v\n", cap(a))

	fmt.Println("***")

	b := []int{1, 2, 3, 4, 5}
	// c := b[:len(a)-1]
	c := append(b[:2], b[3:]...)
	fmt.Println(c)
}
