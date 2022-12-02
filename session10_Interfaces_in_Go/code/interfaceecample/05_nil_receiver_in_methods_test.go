package interfaceecample

import (
	"fmt"
	"testing"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("Got <nil> receiver")
		return
	}
	fmt.Println(t.S)
}

func describeInterface(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TestNilValue(t *testing.T) {
	var i I

	var val *T // Not initialised!
	i = val
	describeInterface(i)
	i.M()   // nil receiver
	val.M() // nil receiver

	// OUTPUT:
	// (<nil>, *interfaceecample.T)
	// Got <nil> receiver
	// Got <nil> receiver
}
