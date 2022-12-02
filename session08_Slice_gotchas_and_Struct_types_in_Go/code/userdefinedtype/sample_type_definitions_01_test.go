package userdefinedtype

import "fmt"

type ID string

type IDs []ID

type IDSet map[ID]struct{}

type Human struct {
	ID   ID
	Name string
	Age  int
}

// Hello is a function with side effects.
func (h Human) Hello() {
	fmt.Printf("Hello, my name is %q, my ID is %q, I'm %d years old", h.Name, h.ID, h.Age)
}
