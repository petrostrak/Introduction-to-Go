package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)    // returns a new bufio.Reader. The Reader will read from standard input (keyboard)
	input, err := reader.ReadString('\n')  // Returns what the user typed, as a string
	fmt.Println("You' ve entered:", input) // Everything up until the newline rune will be read
	log.Fatal(err)                         // Report the error and stop the program
}

/*
Functions and methods like ReadString return an error value of nil, which
basically means “there’s nothing there.” In other words, if err is nil, it means
there was no error.
*/
