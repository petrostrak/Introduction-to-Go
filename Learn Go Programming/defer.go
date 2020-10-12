/*
	Defer
		Used to delay execution of a statement until function exist
		Useful to group "open" and "close" functions together
			Be careful in loops
		Run in LIFO (last in, first out) order
		Arguments evaluated at time defer is executed, not at time of called function execution
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
