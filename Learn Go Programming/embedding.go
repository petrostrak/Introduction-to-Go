package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

type Person struct {
	Name string `Petros`
	Sex  string
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false

	fmt.Println(b)

	// Literal syntax

	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(c)

	a := reflect.TypeOf(Person{})
	field, _ := a.FieldByName("Name")
	fmt.Println(field.Tag)
}
