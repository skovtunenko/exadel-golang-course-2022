package structexample

import (
	"testing"
)

type Point struct {
	X, Y int
}

func TestDumpValues(t *testing.T) {
	var (
		p = Point{1, 2}  // has type Point
		q = &Point{1, 2} // has type *Point
		r = Point{X: 1}  // Y:0 is implicit
		s = Point{}      // X:0 and Y:0
	)

	t.Log(p, q, r, s)
}
