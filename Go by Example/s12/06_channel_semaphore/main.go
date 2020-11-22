/*
A semaphore is a variable (or an abstract data type) used to
control access to a common resource shared by multiple proesses
in a conc urrent system such as a multitasking operation system.

Semaphores are a useful tool on the prevention of a race condition.
However, their use is by no means a guarantee that a programm
is free from these problems.

Semaphores which allow an arbitrary resource count are called counting
semaphores, while semaphores which are restricted to the values 0 and 1
(or locked/ unlocked, unavailable/ available) are called binary semaphores
and are used to implement locks.
*/

package main

import "fmt"

var wordSet1 = []string{"small", "medium", "large"}
var wordSet2 = []string{"red", "greed", "blue", "yellow"}

var c = make(chan string)
var b = make(chan bool)

func main() {
	f := []interface{}{sender1, sender2}
	for i := range f {
		go f[i].(func())() // (func(type))(value)
	}
	go closeSenders()

	fmt.Println("Before getting to the 'channel for range loop'.")
	for val := range c {
		fmt.Println(val)
	}
}

func sender1() {
	for _, w := range wordSet1 {
		c <- w
	}
	b <- true
}

func sender2() {
	for _, w := range wordSet2 {
		c <- w
	}
	b <- true
}

func closeSenders() {
	<-b
	<-b
	close(c)
}
