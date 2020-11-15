package main

import "fmt"

type player struct {
	string
	int
}

func main() {
	p1 := player{"Muhammad Ali", 99} // Positional init
	fmt.Println("Player 1:", p1)
	fmt.Printf("p1.int=%d p1.string=%s", p1.int, p1.string)
}
