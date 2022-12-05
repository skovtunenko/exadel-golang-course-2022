package advancedexample

import (
	"fmt"
	"sync"
	"testing"
)

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	send := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	for _, c := range cs {
		wg.Add(1)
		go send(c)
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

func TestMergeTwoChannelsExample(t *testing.T) {
	in1 := asChan(1, 1, 1)
	in2 := asChan(2, 2, 2)

	for val := range merge(in1, in2) {
		fmt.Println(val)
	}

	fmt.Println("Done.")
}
