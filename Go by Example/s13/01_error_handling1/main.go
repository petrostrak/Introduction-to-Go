package main

import (
	"fmt"
	"log"
	"os"
)

var fileName = "non_existing.txt"

func main() {
	// testCase1()
	// fmt.Println("Control has returned to main after an error in testCase1().")

	// testCase2()
	// fmt.Println("Control does not return to main after an error in testCase2().")

	testCase3()
	fmt.Println("Control does not return to main after an panic in testCase3().")
}

func testCase1() {
	_, err := os.Open(fileName)
	if err != nil {
		log.Println("Error: ", err) // Catches the error and returns to the flow with a timestamp
	}
}

func testCase2() {
	_, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err) // Catches the error but breaks the flow with exit status 1
	}
}

func testCase3() {
	_, err := os.Open(fileName)
	if err != nil {
		panic(err) // Catches the error but breaks the flow with exit status 2
	}
}
