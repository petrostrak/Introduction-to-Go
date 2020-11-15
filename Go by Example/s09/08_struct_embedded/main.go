package main

import "fmt"

type generalInfo struct {
	country, hairColor string
}

type player struct {
	name, sport string
	age         int
	info        generalInfo
}

func main() {
	var p1 player

	p1.name = "Petros Trak"
	p1.age = 34
	p1.sport = "Coding"
	p1.info.country = "Greece"
	p1.info.hairColor = "Dark"

	fmt.Println("P1 :", p1)

	info2 := generalInfo{
		country:   "Japan",
		hairColor: "Blond",
	}

	p2 := player{
		name:  "Eleni Apost",
		age:   28,
		sport: "Japanese",
		info:  info2,
	}

	fmt.Println("P2 :", p2)

	p3 := player{
		name:  "Giannis Lio",
		age:   33,
		sport: "Dota2",
		info: generalInfo{
			country:   "Greece",
			hairColor: "No hair",
		},
	}

	fmt.Println("P3 :", p3)
}
