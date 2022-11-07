package main

import (
	"fmt"
)

// --------------------------------

type oper interface {
	op() string
}

type B struct {
	s string
}

func (a *B) op() string {
	return a.s
}

func interfacePanic(o oper) {
	fmt.Println(o.op() + "!")
	panic("I'm outta here")
}

// --------------------------------

type A struct {
	i int
	s string
}

func (a A) iPanic(b bool) {
	fmt.Println(a.i + 1)
	panic("I'm outta here")
}

func main() {
	// Using interfaces:
	b := &B{"op"}
	interfacePanic(b)

	// Uncomment to see effect:
	// a := A{i: 50}
	// a.iPanic(true)
}
