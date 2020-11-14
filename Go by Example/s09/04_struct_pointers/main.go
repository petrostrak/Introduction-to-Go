package main

import (
	"fmt"
	"time"

	"github.com/fxtlabs/date"
)

type employee struct {
	name, job    string
	lastLoggedIn string
	DOB          date.Date
}

func main() {
	var emp employee

	emp.name = "Nick"
	emp.job = "Go Programmer"
	// Ref: https://golang.org/pkg/time/#pkg-constants
	emp.lastLoggedIn = time.Now().Format(time.RFC850)

	// To use date.Today(), "github.com/fxtlabs/date" needs be installed using:
	// go get github.com/fxtlabs/date
	// Ref: https://godoc.org/github.com/fxtlabs/date
	emp.DOB = date.Today()

	fmt.Println(emp)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	p := &emp

	p.name = "Jack"
	emp.job = "Go Expert"
	fmt.Println(*p)

	fmt.Printf("%x %x", &emp.name, &p.name)
}
