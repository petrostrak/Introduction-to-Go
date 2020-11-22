/*
Race Condition happens when more than one process tries to
access the same resource.
> go run -race main.go (to find possible race conditions)
it prints the results, and gives some information about the race.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitG sync.WaitGroup
var counter = 0

func main() {
	waitG.Add(2)
	go numbers(1)
	go numbers(2)
	waitG.Wait()

	fmt.Printf("\nCounter: %d", counter)
}

func numbers(callID int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10; i++ {
		tmpCounter := counter
		tmpCounter++

		time.Sleep(200 * time.Millisecond)
		counter = tmpCounter

		fmt.Printf("(%d) %d %d\n", callID, rand.Intn(20)+20, counter)
	}
	waitG.Done()
}
