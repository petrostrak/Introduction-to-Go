package main

import "fmt"

type player struct {
	name, sport string
	age         int
}

func main() {
	player1 := player{"Petros Trak", "Soccer", 34} // Posisional approach
	fmt.Println("Player 1=", player1)
	fmt.Printf("(1) name=%s age=%d\n", player1.name, player1.age)

	player2 := player{
		age:   33,
		sport: "Tennis",
		name:  "Greg Bailas",
	}

	fmt.Println("Player 2=", player2)
	fmt.Printf("(2) name=%s age=%d\n", player2.name, player2.age)

	player3 := new(player)
	player3.name = "Giannis Liolis"
	player3.sport = "Dota2"
	player3.age = 33

	fmt.Println("Player 3=", player3)
	fmt.Printf("(3) name=%s age=%d\n", player3.name, player3.age)

	player4 := player{
		name:  "Maggie Trak",
		sport: "Basketball",
	}

	fmt.Println("Player 4=", player4)
	fmt.Printf("(4) name=%s age=%s\n", player4.name, player4.sport)
}
