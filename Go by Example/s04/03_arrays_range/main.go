package main

import (
	"fmt"
)

func main() {
	x := [...]float32{2.1, 3.2, 4.7}

	// x := [3]float32{
	// 	2.1,
	// 	3.2,
	// 	4.7,
	// }

	fmt.Println(x)

	// var total float32 = 0
	var total float32

	//for i, val := range x {	//compiler error - i declared but not used
	for _, val := range x {
		total += val
	}

	fmt.Println(total)
}
