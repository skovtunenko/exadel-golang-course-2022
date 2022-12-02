package errorexample_01

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func A() error {
	return errors.New("cause of the error")
}

func B() error {
	if err := A(); err != nil {
		return errors.WithMessage(err, "unable to do B()")
	}
	return nil
}

func C() error {
	if err := B(); err != nil {
		return errors.WithMessage(err, "unable to do C()")
	}
	return nil
}

func TestError_wrapping(t *testing.T) {
	gotErr := C()
	require.Error(t, gotErr) // change to require.NoError()
}
