package sliceexample

import "fmt"

func ExampleSliceIteration() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// OUTPUT:
	// 2**0 = 1
	// 2**1 = 2
	// 2**2 = 4
	// 2**3 = 8
	// 2**4 = 16
	// 2**5 = 32
	// 2**6 = 64
	// 2**7 = 128
}

