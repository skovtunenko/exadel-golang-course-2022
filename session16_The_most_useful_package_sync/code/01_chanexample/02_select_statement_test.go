package chanexample

import (
	"testing"
)

func f() int {
	return 42
}
func TestSelectStatementMix(t *testing.T) {
	var a []int

	var c1, c2, c3, c4 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		print("received ", i1, " from c1\n")
	case c2 <- i2:
		print("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			print("received ", i3, " from c3\n")
		} else {
			print("c3 is closed\n")
		}
	case a[f()] = <-c4:
	// same as:
	// case t := <-c4
	//	a[f()] = t
	default:
		print("no communication\n")
	}

}
func TestSendRandomSequenceOfBits(t *testing.T) {
	var c = make(chan int)

	for { // send random sequence of bits to c
		select {
		case c <- 0: // note: no statement, no fallthrough, no folding of cases
		case c <- 1:
		}
	}
}

func TestBlockForever(t *testing.T) {
	select {} // block forever
}
