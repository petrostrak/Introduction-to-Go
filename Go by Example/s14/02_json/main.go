/*
JavaScript Object Notation is a standard notation for sending and receiving structured information.package main

Go offers support for encoding / decoding such formats (packages like encoding / json, encoding / xml)

Marshalling is the process of transforming the memory representation of an object to a data format suitable for
storage or transmission, and it is typically used when data must be moved between different parts of a computer
program or from one program to another.

Marshalling is similar to serialization and is used to communicate to remote objects with an object, in this case a serialized
object. It simplifies complex communication (using composite objects, instead of primitives)

The inverse of marshalling is called unmarshalling (or demarshalling, similar to deserialization)
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// SoccerClubs ...
type SoccerClubs struct {
	Name    string
	Country string
	Value   float32 `json:"Dollar Value(B)"`
	Players []string
}

var teams = []SoccerClubs{
	{Name: "Manchester United", Country: "England", Value: 3.69,
		Players: []string{"Eric Cantona", "David Beckham", "Ryan Giggs"}},
	{Name: "Barcelona", Country: "Spain", Value: 3.64,
		Players: []string{"Johan Cruyff", "Xavi", "Andrés Iniesta"}},
	{Name: "Real Madrid", Country: "Spain", Value: 3.58,
		Players: []string{"Zidane", "Roberto Carlos", "Ronaldo"}},
}

func main() {
	fmt.Println("\n- - - - - - - - - - - - - - - - - - - - - - - - - ")

	// use either json.Marshal() or json.MarshalIndent()
	data, err := json.Marshal(teams)
	// data, err := json.MarshalIndent(teams, "", "    ")

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	//fmt.Printf("%v\n\n", teams)
	fmt.Printf("%s\n", data)

	fmt.Println("\n- - - - - - - - - - - - - - - - - - - - - - - - - ")
	var names []struct{ Name string }

	if err := json.Unmarshal(data, &names); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
}
