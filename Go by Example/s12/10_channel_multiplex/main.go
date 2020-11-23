/*
Multiplexing: In telecommunications and computer networks, multiplexing
is a method by which multiple analog or digital signals are combined
into one signal over a shared medium. The aim is to share a scarce resource.

To extend this concept to a typical client-server application, name clients
may send requests to a server (that is the shared scarce resource), requesting
the server to do something for them. The server then processes the
received requests and sends responses back to the clients.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	stats := make(map[string]int)

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	for i := 0; i < 20; i++ {
		go func() {
			time.Sleep(10 * time.Second)
			c1 <- "Hello from customer service #1"
		}()

		go func() {
			time.Sleep(8 * time.Second)
			c2 <- "Hello from customer service #2"
		}()

		go func() {
			time.Sleep(6 * time.Second)
			c3 <- "Hello from customer service #3"
		}()

		select {
		case msg1 := <-c1:
			stats["Petros"]++
			time.Sleep(time.Second)
			fmt.Println(msg1)
		case msg2 := <-c2:
			stats["Giannis"]++
			time.Sleep(time.Second)
			fmt.Println(msg2)
		case msg3 := <-c3:
			stats["Greg"]++
			time.Sleep(time.Second)
			fmt.Println(msg3)
		default:
			stats["Customer is waiting"]++
			time.Sleep(2 * time.Second)
			fmt.Println("No customer service is available at the moment...")
		}
	}
	fmt.Printf("\n***Customer Service Stats***\n%v", stats)
	close(c1)
	close(c2)
	close(c3)
}
