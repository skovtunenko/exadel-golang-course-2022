package deferexample_01

import (
	"fmt"
	"sync"
)

var mx sync.Mutex

func fnUsingDefer() {
	fmt.Println("Preparation...")

	mx.Lock()
	defer mx.Unlock()
	fmt.Println("Inside the critical section...")
	// ..........
	fmt.Println("Still inside the critical section...")
}

func fnWithoutDefer() {
	fmt.Println("Preparation...")

	mx.Lock()
	fmt.Println("I'm doing something...")
	// ..........
	fmt.Println("Still inside the critical section...")
	mx.Unlock()
}

func fnUsingAnonymousFnWithDefer() {
	fmt.Println("Preparation...")

	func() {
		mx.Lock()
		defer mx.Unlock()

		fmt.Println("Inside the critical section...")
	}()

	// ..........
	fmt.Println("Outside of the critical section")
}
