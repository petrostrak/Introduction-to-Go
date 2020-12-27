/*
Reflection is the ability of a program to study its own structure, particularly through the types.

It can be used to dynamically examine types and variables at runtime.

We can use reglection when we need to deal with values of types that don't satisfy a common interface, or they're absent at design time.
fmt.Fprintf is a good example that prints different types and accepts a variable number of parameters.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getType(10))
	fmt.Println(getType(true))
	fmt.Println(getType("hello"))
	fmt.Println(getType(10.31))

}
func getType(t interface{}) string {

	result := ""

	switch t := t.(type) {

	case int: // ...similar cases for int16, ...
		result = strconv.Itoa(t) + "/int"

	case bool:
		if t {
			result = "true/bool"
		} else {
			result = "false/bool"
		}
	case string:
		result = t + "/string"
	default:
		result = fmt.Sprintf("%v/unknown", t)
	}

	return result
}
