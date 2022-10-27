package slicegotchas

import "fmt"

func ExampleCapacityIsNotReached() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	fmt.Println("before:", s, len(s), cap(s)) // before: [1 2 3] 3 4
	reverseSlice(s)
	fmt.Println("after :", s, len(s), cap(s)) // after : [999 3 2] 3 4

	// OUTPUT:
	// before: [1 2 3] 3 4
	// after : [999 3 2] 3 4
}
func reverseSlice(s []int) {
	// The new slice has new length attribute which isn't a pointer, but it still points to the same array.
	// Because the original slice passed “by value” its length wasn’t altered.
	s = append(s, 999)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
