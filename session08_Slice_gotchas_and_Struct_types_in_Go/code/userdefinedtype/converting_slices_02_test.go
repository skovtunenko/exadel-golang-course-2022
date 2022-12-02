package userdefinedtype

import "fmt"

func ExampleFailedCompilation() {
	type T1 int
	type T2 int
	var t1 T1

	var x = T2(t1) // OK
	_ = x

	var src []T1
	_ = src
	// var sx = ([]T2)(src) // NOT OK, code will not compile!
}

func ExampleProperConversion() {
	type T1 int
	type T2 int

	src := []T1{1, 2, 3, 4}
	fmt.Println(src)

	dest := make([]T2, 0, len(src))
	for _, val := range src {
		dest = append(dest, T2(val))
	}

	fmt.Println(dest)

	// Output:
	// [1 2 3 4]
	// [1 2 3 4]
}
