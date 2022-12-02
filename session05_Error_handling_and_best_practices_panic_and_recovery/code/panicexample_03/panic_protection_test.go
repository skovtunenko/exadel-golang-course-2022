package panicexample_03

import (
	"log"
	"runtime/debug"
	"testing"
)

func TestPanicProtection(t *testing.T) {
	// Case #1:
	// panicingFn()

	// Case #2:
	protect(panicingFn)
}

func panicingFn() {
	panic("BOOM!")
}

func protect(fn func()) {
	defer func() {
		log.Println("this will be printed normally, even if there is a panic / or no panic")
		if x := recover(); x != nil {
			// This is how to print stacktrace from the panic:
			log.Printf("Runtime panic, recovered from: %v; stack: %s", x, debug.Stack())
		}
	}()
	log.Println("start executing fn()...")
	fn()
	log.Println("...fn() executed without panic")
}
