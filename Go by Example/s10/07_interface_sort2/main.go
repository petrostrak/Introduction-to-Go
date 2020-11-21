/*
Type Interface: a type, typically a collection, that satisfies
sort.Interface can be sorted by the routines in the 'sort'
package. The methods require that the elements of the collection
be enumerated by an integer index.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 2, 10, -4, 13, 9, 17}
	fmt.Println(n)

	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

	s := []string{"Petros", "Eleni", "Maggie", "Maria", "Greg", "Giannis"}
	fmt.Println(s)

	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
