package slicegotchas

import "fmt"

func ExampleNewArrayAllocated() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	fmt.Println("before:", s, len(s), cap(s)) // before: [1 2 3] 3 4
	reverseAndAppend(s)
	fmt.Println("after :", s, len(s), cap(s)) // after : [1 2 3] 3 4

	// OUTPUT:
	// before: [1 2 3] 3 4
	// after : [1 2 3] 3 4
}
func reverseAndAppend(s []int) {
	s = append(s, 999, 1000, 1001) // new array will be allocated and then reversed
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
