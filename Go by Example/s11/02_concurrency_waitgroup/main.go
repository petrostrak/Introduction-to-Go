package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitG sync.WaitGroup

func main() {
	waitG.Add(2) // Determines the number of go-routines that will be watched in ths programm
	go numbers()
	go alphabets()
	waitG.Wait() // Blocks until all go-routines are finished

	time.Sleep(3200 * time.Millisecond)
	fmt.Println("\nMain terminated")
}

func numbers() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 15; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", rand.Intn(20)+20)
	}
	waitG.Done()
}

func alphabets() {
	for i := 'C'; i <= 'G'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	waitG.Done()
}
