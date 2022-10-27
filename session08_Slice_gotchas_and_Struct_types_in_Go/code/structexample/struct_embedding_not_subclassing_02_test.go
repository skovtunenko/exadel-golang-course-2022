package structexample

import (
	"fmt"
)

type Human struct {
	Name string
}

func (b Human) Greet() {
	fmt.Println("Hello, i'm", b.Name)
}

type Contractor struct {
	Human
	Name string // This field has same name as in `Human` struct!
}

func ExampleEmbeddingIsNotSubclassing() {
	fb := Contractor{
		Human: Human{Name: "Ivan"},
		Name:  "Petr",
	}
	fb.Greet() // call will be promoted to `fb.Human` field, which doesn't know anything about `Contractor.Name` field!
	fb.Human.Greet()

	// OUTPUT:
	// Hello, i'm Ivan
	// Hello, i'm Ivan
}
