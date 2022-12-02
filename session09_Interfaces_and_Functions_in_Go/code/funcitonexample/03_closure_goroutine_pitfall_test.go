package funcitonexample

import (
	"fmt"
	"testing"
)

func TestBrokenGoroutinePitfall(t *testing.T) {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println("value:", v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

	// OUTPUT:
	// value: c
	// value: c
	// value: c
}

func TestFixedGoroutinePitfall(t *testing.T) {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v // create a new local variable `v`
		go func() {
			fmt.Println("value:", v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

	// OUTPUT might be random, but will contain all three letters:
	// value: a
	// value: c
	// value: b
}
