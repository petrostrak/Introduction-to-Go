/*
Buffered Channels (asynchronous/ non-blocking):

The capacity argument of make is bigger than zero. c := make(chan int, 3)
c := make(chan channelType, channelCapacity)

a send operation: adds an element ot the end of the queue
a receive operation: removes an element from the front

'sends' to a buffered channel are blocked only when the buffer is full
if the channel is 'full' ->
'send' operation blocks until a 'receive' operation makes a space available

'receives' from a buffered channel are blocked only when the buffer is empty
If the channel is 'empty' ->
'receive' operation blocks until a 'send' operation send a value
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)

	for i := 1; i <= 3; i++ {
		go printMsg(c, i)
	}
	time.Sleep(10 * time.Second)
}

func printMsg(c chan int, id int) {
	fmt.Printf("ooo %d is waiting for a channel space...\n", id)

	c <- id
	fmt.Printf("=== %d has a channel space\n", id)

	time.Sleep(600 * time.Millisecond)
	fmt.Printf("xxx %d has released the channel space\n", id)

	<-c
}
