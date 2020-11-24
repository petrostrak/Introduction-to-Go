// ASSIGNMENT: Rewrite the previous assignment (pipeline-like)
// using functions and unidirectional Channels.

package main

import (
	"fmt"
	"strings"

	tjarratt "github.com/tjarratt/babble"
)

func main() {
	newWords := make(chan string)
	uWords := make(chan string)

	go sendWords(newWords)
	go converWords(uWords, newWords)
	printWords(uWords)
}

func sendWords(out chan<- string) {
	babbler := tjarratt.NewBabbler()
	for i := 0; i < 5; i++ {
		out <- babbler.Babble() // Produces random words
	}
	close(out)
}

func converWords(out chan<- string, in <-chan string) {
	for w := range in {
		out <- w + " --> " + strings.ToUpper(w)
	}
	close(out)
}

func printWords(in <-chan string) {
	for w := range in {
		fmt.Println(w)
	}
}
