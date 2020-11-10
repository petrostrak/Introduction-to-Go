package main

import "fmt"

func main() {
	add := func(nums ...int) int { // variadic
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	fmt.Println(add(1, 2, 3, 4))

	nums := []int{1, 2, 3, 4}
	fmt.Println(add(nums...))

	fmt.Printf("%v\n", calc(add, nums))

	mult := func(nums ...int) int {
		total := 1
		for _, num := range nums {
			total *= num
		}
		return total
	}

	fmt.Printf("%v\n", calc(mult, nums))
}

func calc(f func(...int) int, x []int) int {
	return f(x...)
}
