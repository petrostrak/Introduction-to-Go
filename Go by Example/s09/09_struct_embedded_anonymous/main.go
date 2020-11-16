package main

import "fmt"

type generalInfo struct {
	country, hairColor string
}

type player struct {
	name, sport string
	age         int
	generalInfo
}

func main() {
	var p1 player
	p1.age = 34
	p1.name = "Petros Trak"
	p1.sport = "Soccer"

	p1.country = "Greece"
	p1.hairColor = "Brown"

	fmt.Println(p1)
}
