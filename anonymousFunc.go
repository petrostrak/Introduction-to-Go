package main

import (
	"fmt"
)

func main() {
	func() {
		msg := "Hello GO!"
		fmt.Println(msg)
	}()

	// // for asynchronous functions
	// for i := 0; i < 5; i ++ {
	//   func (i int)  {
	//     fmt.Println(i)
	//   }(i)
	// }
}
