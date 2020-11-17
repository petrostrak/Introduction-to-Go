package main

import (
	"fmt"
	"strings"
)

type movie struct {
	name  string
	actor string
}

type imdb struct {
	name    string
	comment string
	movie
}

func (m movie) fullInfo() string {
	return m.name + "-" + m.actor
}

func (i imdb) fullInfo() string {
	return strings.ToLower(i.movie.name + "-" +
		i.movie.actor + "-" + i.name + "-" + i.comment)
}

func main() {
	m := movie{"Vampire's kiss", "Nicolas Cage"}
	Imdb1 := imdb{"Sony Pictures", "Awesome Movie", m}

	Imdb2 := imdb{
		movie: movie{
			name:  "The Godfather",
			actor: "Marlon Brando",
		},
		name:    "Paramount Pictures",
		comment: "Superb!",
	}

	fmt.Println(m.fullInfo())
	fmt.Println(Imdb1.fullInfo())
	fmt.Println(Imdb2.fullInfo())

	fmt.Printf("%s-%s \n", Imdb2.movie.name, Imdb2.name)
}
