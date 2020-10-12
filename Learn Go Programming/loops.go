/*
  For statements
    Simple loops
      for initializer; test; incrementer {}
      for test {}
      for {}
    Exiting early
      break
      continue
      labels (Loop: break Loop)
*/

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//function-scoped for loop
	k := 0
	for ; k < 5; k++ {
		fmt.Println(k)
	}

	// or..
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// continue
	for a := 0; a < 5; a++ {
		if a%2 == 0 {
			continue
		}
		fmt.Println(a)
	}

	// nested loop
Loop:
	for b := 0; b < 5; b++ {
		for c := 0; c < 5; c++ {
			fmt.Println(b * c)
			if b*c >= 3 {
				break Loop
			}
		}
	}

	// loop through an array
	z := []int{1, 2, 3}
	for k, v := range z {
		fmt.Println(k, v)
	}

	y := "Hello Go!"
	for k, v := range y {
		fmt.Println(k, string(v))
	}

	f := "If I dont want the index!"
	for _, v := range f {
		fmt.Println(string(v))
	}
}
