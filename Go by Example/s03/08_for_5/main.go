package main

import "fmt"

func main() {
	fmt.Println("Press CTRL+C to stop the program.")

	i := 2
	for {
		fmt.Print(i, " ")
		i++
	}
}
