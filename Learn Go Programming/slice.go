/*
	Slices
		Backed by array
		Creation styles
      Slice existing array of Slice
      Literal styles
      Via make function
        a := make([]int, 10) // create slice with capacity and length 10
        a := make([]int, 10, 100) // create slice with length 10 and capacity 100
		len function returns length of arrays
    cap function returns length of underlying array
	 	append function to add elements to slice
      May cause expensive copy operation if underlying array is too small
    Copies refer to same underlying array
*/

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // The syntax of slice is the exact same with arrays but without the [...] at declaration
	b := a[:]                                 // slice all of the elements of a
	c := a[3:]                                // slice from 4th element towards
	d := a[:6]                                // slice first 6 elements
	e := a[3:6]                               // slice the 4th, 5th and 6th elements
	a[5] = 99                                 // it'll affect all cause they are pointing to the same underline array
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
