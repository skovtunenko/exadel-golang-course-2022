package methodexample

import (
	"fmt"
	"testing"
)

type Human struct {
	name string
}

func (h Human) getName() string {
	return h.name
}

func (h *Human) setName(name string) {
	h.name = name
}

func (h *Human) String() string {
	return fmt.Sprintf("Human=%v", *h) // change here `*h` to just `h`
}

func TestMethodsAreJustFunctions(t *testing.T) {
	human := Human{name: "Sergio"}

	// Assign method on VALUE RECEIVER to the variable:
	var getFn = Human.getName
	gotName := getFn(human)
	t.Logf("got name: %q\n", gotName)

	// Assign method on POINTER RECEIVER to the variable:
	var setFn = (*Human).setName
	setFn(&human, "new name")
	t.Logf("After chaning the name: %q\n", human.name)
}

func TestStackOverflow(t *testing.T) {
	human := Human{name: "Sergio"}

	t.Log(human.String())
}
