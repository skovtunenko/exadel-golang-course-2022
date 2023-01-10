package goroutinesexample

import (
	"fmt"
	"testing"
)

func TestBestPractice1(t *testing.T) {
	c := make(chan string)

	for i := 1; i <= 2; i++ {
		go func(i int, co chan<- string) {
			for j := 1; j <= 2; j++ {
				co <- fmt.Sprintf("hi from %d.%d", i, j)
			}
		}(i, c)
	}

	for i := 1; i <= 4; i++ {
		fmt.Println(<-c)
	}
}

// OUTPUT:
// Hi from 2.1
// Hi from 2.2
// Hi from 1.1
// Hi from 1.2
