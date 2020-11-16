package main

import "fmt"

// ASSIGNMENT: Multiple embedded structs - Write a program to demonstrate
// multiple embedded structs. Something like (info -> person -> student)
// Similar to inheritance.

type info struct {
	telephone string
	email     string
	address   string
}

type person struct {
	firstName string
	lastNale  string
	age       int
	info
}

type student struct {
	fieldOfStudy string
	yearOfStudy  int
	isStudent    bool
	person       person
}

func main() {
	person1 := person{"Petros", "Trak", 34, info{"6942146041", "petros@gmail.com", "Munster St. 91"}}
	fmt.Println(person1)

	var student1 student
	student1.fieldOfStudy = "Mechanical Engineering"
	student1.yearOfStudy = 4
	student1.isStudent = false
	student1.person = person1

	fmt.Println(student1)
}
