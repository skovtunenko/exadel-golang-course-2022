package interfacesexample

import "testing"

// Values of interface `GetSet` can be assigned to values of type `Getter`.
type Getter interface {
	Get() string
}

// Values of interface `GetSet` can be assigned to values of type `Setter`.
type Setter interface {
	Set(val string) // same as just: Set(int)
}

// Values of interfaces `Getter` or `Setter` CAN NOT be assigned to values of type `GetSet`.
type GetSet interface {
	Getter
	Setter
}

type Human struct {
	name string
}

func (h *Human) Get() string {
	if h == nil {
		return "<Unknown name>"
	}
	return h.name
}

func (h *Human) Set(name string) {
	if h == nil {
		return
	}
	h.name = name
}

func TestHuman(t *testing.T) {
	h := &Human{name: "Sergio"}

	var getter Getter = h
	t.Logf("calling getter.Get(): %q\n", getter.Get())

	var setter Setter = h
	setter.Set("Another name")

	t.Logf("again calling getter.Get(): %q\n", getter.Get())
}
