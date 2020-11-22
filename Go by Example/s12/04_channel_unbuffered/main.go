/*
By default, communication is unbuffered (synchronous)

It means that a 'send' does not complete until a 'receiver'
eccepts  the value.

This blocking behaviour of unbuffered channels tells us that
there is no space in the channel for data (because for any 'send'
there must be a 'receive' and vise versa)

In short, channel operations (i.e. send/receive) block until
the other side is ready.

Communication can be seen as a form of synchronization when
go-routines share data through a channel that synchronizes(takes turn)
to communicate with those go-routines. Clearly, unbuffered channels
are the perfect candidate for synchronizing communication of multiple
go-routines.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var waitG sync.WaitGroup

func main() {
	c := make(chan int)

	waitG.Add(1)
	go send(c)
	go receive(c)
	waitG.Wait()

	time.Sleep(2 * time.Second)
}

func send(c chan int) {
	for i := 1; i < 6; i++ {
		c <- i
	}
	waitG.Done()
}

func receive(c chan int) {
	for {
		fmt.Println(<-c)
	}
	// No waitG.Done() as it won't be reached
}
