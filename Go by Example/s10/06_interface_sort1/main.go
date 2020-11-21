/*
Package sort provides primitives for sorting slices  and
user-defined collections
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 2, 10, -4, 13, 9, 17}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	s := []string{"Petros", "Eleni", "Maggie", "Maria", "Greg", "Giannis"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}
