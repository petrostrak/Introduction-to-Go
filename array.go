/*
	Arrays
		Collection of items with same type
		Fixed size
		Declaration styles
			a := [3]int{1, 2, 3}
			a:= [...]int{1, 2, 3}
			var a [3]int
		Access via zero-based index
			a := [3]int{1, 3, 5} // a[3] == 3
		len function returns size of arrays
		Copies refer to different underlying data
*/

package main

import (
	"fmt"
)

func main() {
	grades := [...]int{97, 68, 35}
	fmt.Printf("Grades : %v\n", grades)

	fmt.Println("***")

	a := [...]int{1, 2, 3}
	b := a // Assigning an variable to an array, will not point to it, instear it will copy the array.
	// b := &a // To point to an array without coping, we use the & operator
	b[1] = 5
	fmt.Println("Array a :", a)
	fmt.Println("Array b :", b, " after coping (declaring) b = a.")

	fmt.Println("***")

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Printf("Identity Matrix : %v\n", identityMatrix)
}
