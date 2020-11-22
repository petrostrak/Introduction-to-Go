/*
All channels are created bidirectional, but we can assign
them to directional channel variables.

func bi(c chan int) {	// a "bidirectional" channel
	c <- 10				// ok
	fmt.Println(<-c)	// ok
}

func snd(out chan<- int) {	// a "send-only" channel
	out <- 10				// ok
	// fmt.Println(<-out)	// error - receive from send-only type
}

func rcv(in <-chan int) {	// a "receive-only" channel
	// in <- 10 			// error - send to receive-only type
	fmt.Println(<-in)		// ok
}

Receive-only channels cannot be closed, because a channel is a
way for the sender to say that no more values will be sent to the
channel, that doesn't make sense for receive-only channels.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ASSIGNMENT: Assuming your current bank balance is 200, write
// a program containing these functions:
// 1) credit() function with a 'bi-directional channel' parameter that sends random
// numbers (between 1 and 10) to the channel that will be added to the balance.
// 2) debit() function with a 'send-only channel' parameter that sends random
// numbers (between -10 and -1) to the channel that will be subtracted from the balance.
// 3) balance() function with a 'receive-only channel' parameter that adds/subtracts
// the random values (generated in steps 1 & 2) to the balance.
//
// Here's a sample output that indicates that the balance is randomly changing:
//		Press Enter to stop the program ...
// 		=> 200 + (6) = 206
// 		=> 206 + (-10) = 196
// 		=> 196 + (3) = 199
// 		=> 199 + (-2) = 197
//

var currentBalance = 200

func main() {
	fmt.Println("Press Enter to stop the program...")
	rand.Seed(time.Now().UnixNano())

	c := make(chan int) //bidirectional

	go credit(c)
	go debit(c)
	go balance(c)

	var input string
	fmt.Scanln(&input) // press "Enter" to exit
}

func credit(c chan int) { // a bidirectional channel
	for i := 0; ; i++ {
		c <- rand.Intn(9) + 1 // random(1, 10)
	}
}

func debit(c chan<- int) { // a send-only channel
	for i := 0; ; i++ {
		c <- rand.Intn(9) - 10 // random(-10, -1)
	}
}

func balance(c <-chan int) { // a receive-only channel
	for {
		num, ok := <-c
		if ok == false {
			fmt.Println("Error")
			break
		}

		oldBalance := currentBalance
		currentBalance += num

		fmt.Printf("=> %d + (%d) = %d\n", oldBalance, num, currentBalance)
		time.Sleep(1 * time.Second)
	}
}
