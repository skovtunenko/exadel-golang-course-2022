package slicegotchas

import "fmt"

func ExamplePassSliceToFunction() {
	var s []int // nil slice
	for i := 1; i <= 3; i++ {
		s = append(s, i) // append() works fine with nil slices
	}
	fmt.Println("before:", s, len(s), cap(s)) // before: [1 2 3] 3 4
	reverseTheSlice(s)
	fmt.Println("after: ", s, len(s), cap(s)) // after:  [3 2 1] 3 4

	// OUTPUT:
	// before: [1 2 3] 3 4
	// after:  [3 2 1] 3 4
}

// function to reverse the content of the slice
func reverseTheSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
