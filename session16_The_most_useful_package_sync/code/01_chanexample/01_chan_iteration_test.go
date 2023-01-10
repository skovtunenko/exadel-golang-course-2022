package chanexample

import (
	"fmt"
	"testing"
)

func FibonacciProducer(ch chan int, count int) {
	defer close(ch)
	n2, n1 := 0, 1
	for count >= 0 {
		ch <- n2
		count--
		n2, n1 = n1, n2+n1
	}
}

func TestIterationUsingRangeOverChan(t *testing.T) {
	ch := make(chan int)
	go FibonacciProducer(ch, 10)
	idx := 0
	// To break such iteration channel needs to be closed explicitly.
	// Otherwise range would block forever in the same way as for nil channel.
	for num := range ch {
		fmt.Printf("F(%d): \t%d\n", idx, num)
		idx++
	}
}
