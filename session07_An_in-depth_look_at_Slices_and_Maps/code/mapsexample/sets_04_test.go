package mapsexample

import "fmt"

func ExampleSet() {
	setOfStrings := map[string]struct{}{}

	// Add the same values multiple times:
	setOfStrings["one"] = struct{}{}
	setOfStrings["one"] = struct{}{}
	setOfStrings["one"] = struct{}{}

	for k := range setOfStrings {
		fmt.Println("value =", k)
	}

	// output:
	// value = one
}
