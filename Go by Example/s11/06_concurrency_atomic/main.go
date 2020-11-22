/*
Package atomic provides low-level atomic memory primitives
useful for implementing synchronization algorithms.
These functions require great care to be used correctly.
Except for special, low-level applications, synchronization
is better done with channels or the facilities of the sync
package.

Another way of encapsulation an operation:
	func AddUint64(addr *uint64, delta uint64)(new uint64)
AddUint64() atomically adds delta to *addr and returns the
new value
	func LoadUint64(addr *uint64)(val uint64)
LoadUint64() atomically loads *addr
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var waitG sync.WaitGroup
var counter uint64

func main() {
	waitG.Add(2)
	go numbers(1)
	go numbers(2)
	waitG.Wait()

	fmt.Printf("\nCounter: %d", counter)
}

func numbers(callID int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)

		atomic.AddUint64(&counter, 1)
		fmt.Printf("(%d) %d %d\n", callID, randInt(20)+20), atomic.LoadUint64(&counter))
	}
	waitG.Done()
}
