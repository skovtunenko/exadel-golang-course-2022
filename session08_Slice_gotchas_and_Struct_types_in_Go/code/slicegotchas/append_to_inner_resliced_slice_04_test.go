package slicegotchas

import "fmt"

func ExampleAppenToInnerReslicedSlice() {
	array := []string{"a", "b", "c", "d", "e", "f"}

	slice1 := array[:3] // Append operations on this 'slice1' is DANGEROUS!!!
	slice2 := array[3:]

	fmt.Printf("so far so good: slice1 %v slice2 %v\n", slice1, slice2)

	slice1 = append(slice1, "BANG")

	fmt.Printf("append to slice1: %v\n", slice1)
	fmt.Printf("slice2 is now corrupt: %v\n", slice2)

	fmt.Println("Base array:", array)

	// OUTPUT:
	// so far so good: slice1 [a b c] slice2 [d e f]
	// append to slice1: [a b c BANG]
	// slice2 is now corrupt: [BANG e f]
	// Base array: [a b c BANG e f]
}
