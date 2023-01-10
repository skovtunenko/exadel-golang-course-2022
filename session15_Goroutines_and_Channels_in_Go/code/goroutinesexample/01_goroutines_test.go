package goroutinesexample

import (
	"fmt"
	"testing"
	"time"
)

func printValuesFrom(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func TestGoroutines_01(t *testing.T) {
	printValuesFrom("direct")

	go printValuesFrom("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
