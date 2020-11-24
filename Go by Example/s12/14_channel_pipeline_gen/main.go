// ASSIGNMENT: A pipeline connects two (or more) channels operations, so that the
// result/output of the first channel operations becomes the input of the second channel.
//
// Write a program containing two functions to simulate a pipeline of two channels:
// step 1: function gen(numbers) sends some numbers to channel c.
// step 2: function sq(channel) reads those values from channel c and squares them.
// Ref: https://blog.golang.org/pipelines
//

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	c := gen(nums...)
	out := sq(c)

	for range nums {
		fmt.Printf("%4d ", <-out)
	}

	fmt.Println()
	for n := range sq(sq(gen(nums...))) {
		fmt.Printf("%4d ", n)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
