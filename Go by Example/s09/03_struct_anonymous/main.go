package main

import "fmt"

func main() {

	player1 := struct {
		name, sport string
		age         int
	}{"Petros Trak", "Soccer", 34}

	fmt.Println("Player 1:", player1)

	player2 := struct {
		name, sport string
		age         int
	}{
		age:   34,
		name:  "Giannis",
		sport: "Dota2",
	}

	fmt.Println("Player 2:", player2)
}
