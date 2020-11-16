package main

import (
	"fmt"
	"strings"

	"./types"
)

func changeAthleteName1(p athletes.Player) {
	p.Name = "Petros Trak"
}

func changeAthleteName2(p *athletes.Player) {
	p.Name = "Petros Trak"
	p.Country = strings.ToUpper(p.Country)
}

func main() {
	p1 := athletes.Player{"Maggie Trak", "Housewife", 36, athletes.Info{"Greece", "Brown"}}
	fmt.Println("(1) p1:", p1)

	changeAthleteName1(p1)
	fmt.Println("(2) p1:", p1)

	changeAthleteName2(&p1)
	fmt.Println("(3) p1:", p1)

	fmt.Println("(4) p1:", *p1.ToLowerCase())
}
