package interfaceecample

import (
	"errors"
	"testing"
)

type MyError struct {
}

func (m MyError) Error() string {
	return "my error"
}

func incorrectlyReturnsError() error {
	var p *MyError = nil

	// ..... some code .........

	return p // Will always return a non-nil error.
}

func correctlyReturnsError() error {
	if true {
		return errors.New("some error")
	}
	return nil
}

func TestReturnError(t *testing.T) {
	gotErr := incorrectlyReturnsError()
	t.Logf("got error: %v ; Is this error equal to nil? %t\n", gotErr, gotErr == nil)

	gotErr = correctlyReturnsError()
	t.Logf("got another error: %v\n", gotErr)
}
