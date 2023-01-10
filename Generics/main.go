package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// sumInts adds together the values of m.
func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

// sumInts adds together the values of m.
func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

type Number interface {
	int64 | float64
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func sumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func sumOfArray[V Number](arr []V) V {
	var s V
	for _, v := range arr {
		s += v
	}

	return s
}

type Point []int32

// scale returns a copy of s with each element multiplied by c.
func scale[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}

	return r
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// Initialize an array of int64
	intArr := []int64{2, 4, 6, 8}

	// Initialize an array of float64
	intFloat := []float64{23.43, 55.92, 11.32, 84.60}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		sumInts(ints),
		sumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		sumIntsOrFloats(ints),
		sumIntsOrFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		sumOfArray(intArr),
		sumOfArray(intFloat))

	arr := []int8{1, 3, 5, 6, 9}

	fmt.Printf("Generic scale: %v\n",
		scale(arr, 2))
}
