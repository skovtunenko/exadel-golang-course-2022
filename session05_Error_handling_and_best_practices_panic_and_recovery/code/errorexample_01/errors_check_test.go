package errorexample_01

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type MyError struct {
}

func (m MyError) Error() string {
	return "my error"
}

var ErrBad = &MyError{}

func incorrectlyReturnsError() error {
	var p *MyError = nil
	if false {
		p = ErrBad
	}
	return p // Will always return a non-nil error.
}

func correctlyReturnsError() error {
	if true {
		return ErrBad
	}
	return nil
}

func TestReturnError(t *testing.T) {
	gotErr := incorrectlyReturnsError()
	require.NoError(t, gotErr) // this will fail

	gotErr = correctlyReturnsError()
	require.Error(t, gotErr)
}
