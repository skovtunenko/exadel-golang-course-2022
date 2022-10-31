package funcitonexample

import (
	"testing"
)

type producer func() uint

func makeEvenGenerator() producer {
	i := uint(0) // this variable captured by reference!!!

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func TestGenerator(t *testing.T) {
	nextEvenFn := makeEvenGenerator()

	evenNumber := nextEvenFn()
	t.Log("first call:", evenNumber) // Out: 0
	if evenNumber != 0 {
		t.FailNow()
	}

	evenNumber = nextEvenFn()
	t.Log("second call:", evenNumber) // Out: 2
	if evenNumber != 2 {
		t.FailNow()
	}

	evenNumber = nextEvenFn()
	t.Log("third call:", evenNumber) // Out: 4
	if evenNumber != 4 {
		t.FailNow()
	}
}
