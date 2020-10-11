package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Printf("\ni=%d ** ", i)
		for j := 1; j < 3; j++ {
			fmt.Printf("%d ", i*j)
		}
	}
}
