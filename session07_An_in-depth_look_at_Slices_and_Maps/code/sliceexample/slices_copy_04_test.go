package sliceexample

func ExampleCopy() {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	var b = make([]byte, 5)

	n1 := copy(s, a[0:])           // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
	n2 := copy(s, s[2:])           // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
	n3 := copy(b, "Hello, World!") // n3 == 5, b == []byte("Hello")

	_ = n1
	_ = n2
	_ = n3
}
