package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := -5

	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("Got it!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
}
