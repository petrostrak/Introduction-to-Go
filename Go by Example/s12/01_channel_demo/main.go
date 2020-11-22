/*
Channels connect go-routines and enable thme to communicate
and synchronize their execution

CHannels are typed, means that the data that is passed into
a channel (shared between multiple go-routines) must be of
the same type.

To create a channel:
# c = make(chan int)		// unbuffered channel - c has type 'chan int'
# c = make(chan int, 4)	// buffered channel with capacity 4

Channels are reference types. The zero (default) value of a channel
is nil.

You can compare two channels (of the same type) using ==.
True means they are referring to the same channel variable.
A channel may also be compared to nil.

Two operations of a channel: send and receive (collectively know as
communications).
# c <- x	// send - sends value x to channel c
# x = <-c	// receive - reads from the channel c and saves in variable x
# <-c		// receive - result is discarded

Channels suppert a third operation, close().
close() sets a flag indication that no more values will be sent on this channel.
To close a channel: close(c)

On a closed channel:
# Further 'send' operations will panic.
# Further 'receive' operations receive the already sent values, into no more values are left.
# After the completion point, any receive operations immediately concludes and
# yields the zero value of the channel's element type.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan string)

	start := time.Now()

	go message1(c)
	go message2(c)
	go message3(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	close(c)
	elapsed := time.Since(start)
	fmt.Printf("\nTime elapsed: %s \n", elapsed)
}

func message1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "msg1, delay 3 secs"
}

func message2(c chan string) {
	time.Sleep(4 * time.Second)
	c <- "msg2, delay 4 secs"
}

func message3(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "msg3, delay 2 secs"
}
