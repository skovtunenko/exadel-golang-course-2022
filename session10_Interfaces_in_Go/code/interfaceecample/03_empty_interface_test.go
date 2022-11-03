package interfaceecample

import (
	"fmt"
	"testing"
)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TestEmptyInterface(t *testing.T) {
	var i interface{}
	describe(i) // (<nil>, <nil>)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
}
