/*
Strucs
  Collections of disparate data types that describe a single concept
  Keyed by named fields
  Normally created as types, but anonymous strucs are allowed
  Structs are value types
  No inheritance, but can use composition via Embedding
  Tags can be added to struct fields to describe field
*/
package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[1])
	fmt.Println(aDoctor.actorName)

	bDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(bDoctor)
}
