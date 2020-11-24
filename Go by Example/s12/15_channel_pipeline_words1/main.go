// ASSIGNMENT: Write a program including the following goroutines:
// 1) The 'newWords' goroutine that produces random words and send them to a channel;
// 2) The 'uWords' (uppercase words) goroutine that converts the randomly-generated
// words to uppercase;
// Use anonymous functions to write your goroutines.
//

package main

import (
	"fmt"
	"strings"

	tjarraat "github.com/tjarratt/babble"
)

func main() {
	babbler := tjarraat.NewBabbler()
	newWords := make(chan string)
	uWords := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			newWords <- babbler.Babble() // Produces random word
		}
		close(newWords)
	}()

	go func() {
		for w := range newWords {
			uWords <- w + " --> " + strings.ToUpper(w)
		}
		close(uWords)
	}()

	for w := range uWords {
		fmt.Println(w)
	}
}
