package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go numbers()
	go alphabets()

	time.Sleep(3200 * time.Millisecond)
	fmt.Println("\nMain terminated")
}

func numbers() {
	// func Now() Time
	// returns the current local time. Requires time package
	// func (t time) UnixNano() int64
	// returns 't' as a Unix time, the number of nanoseconds
	// elapsed since 1 Jan 1970 UTC. Requires time package

	rand.Seed(time.Now().UnixNano())

	for i := 1; i < 15; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", rand.Intn(20)+20)
	}
}

func alphabets() {
	for i := 'C'; i < 'G'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
