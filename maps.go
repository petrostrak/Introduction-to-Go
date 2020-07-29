/*
  Maps
    Collections of value types that are accessed via keys
    Created via literals or via make()
    Members accessed via [key] syntax
      myMap["key"] = "value"
    Check for presence with "value, ok" form of result
    Multiple assignments refer to same underlying data
*/

package main

import (
	"fmt"
)

func main() {
	// statePopulation := make(map[string]int)
	// statePopulation = map[string]int{
	statePopulation := map[string]int{
		"California": 39260012,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
		"Ohio":       11614373,
	}
	fmt.Println(statePopulation)
	statePopulation["Georgia"] = 10310371
	fmt.Println(statePopulation)
	delete(statePopulation, "New York")
	fmt.Println(statePopulation)
	fmt.Println(statePopulation["New York"])
	_, ok := statePopulation["New York"]
	fmt.Println(ok)
	_, ok = statePopulation["Ohio"]
	fmt.Println(ok)
	fmt.Println(len(statePopulation))
}
