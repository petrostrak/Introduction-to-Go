/*
Type assertion: detecting and converting the type of an
interface variable.

The following interface type variable 'message' can contain
a value of any type. We need a way to detect this dynamic
type. We will use variable (TYPE)

var message interface{} = 10		or
var message interface{} = "Hello"
s, ok := message.(string)			if ok {...}
*/

package main

import "fmt"

func main() {
	var message interface{} = "Hello!"
	s, ok := message.(string)

	if ok {
		fmt.Printf("Assertion completed successfully! - %q %T \n", s, message)
	} else {
		fmt.Printf("Value is not a string - %q %T \n", s, message)
	}

	f := 20.355
	fmt.Printf("%d %T %T \n", int(f), int(f), f)

	var num interface{} = 10
	println(num.(int) + 20)
	fmt.Printf("%d %T \n", num, num)
}
