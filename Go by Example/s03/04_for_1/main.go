package main

import "fmt"

func main() {
	// to repeat the same operations
	fmt.Println(2, 20)
	fmt.Println(3, 30)
	fmt.Println(4, 40)

	fmt.Println()
	i := 2
	for i < 5 {
		fmt.Println(i, i*10)
		i++
	}
}
