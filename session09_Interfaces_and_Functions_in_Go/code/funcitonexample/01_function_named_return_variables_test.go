package funcitonexample

import (
	"errors"
	"fmt"
	"testing"
)

func TestNamedReturnVariables(t *testing.T) {
	gotErr := functionWithNamedReturnVariables()
	if gotErr == nil {
		t.FailNow()
	}
}

func functionWithNamedReturnVariables() (err error) {
	// If there is a panic we need to recover in a deferred func:
	defer func() {
		if r := recover(); r != nil {
			// the named return value will be modified:
			err = errors.New("return value will be modified successfully because of named return value")
		}
	}()

	legacyFunctionThatWillPanic() // always panic inside

	fmt.Println("This is unreachable call!")
	return err // because of panic, this return statement won't be executed
}

func legacyFunctionThatWillPanic() {
	panic("PANIC!")
}
