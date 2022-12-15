package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNodeStack_E2E(t *testing.T) {
	t.Run("operations_on_nil_stack", func(t *testing.T) {
		r := require.New(t)
		const defaultValue = 0

		var ns *NodeStack[int]

		empty := ns.IsEmpty()
		r.True(empty)

		val, ok := ns.Top()
		r.False(ok)
		r.Equal(defaultValue, val)

		val, ok = ns.Pop()
		r.False(ok)
		r.Equal(defaultValue, val)

		ns.Push(777) // no effect, because the `ns` is nil
		val, ok = ns.Pop()
		r.False(ok)
		r.Equal(defaultValue, val)
	})

	t.Run("popping_from_empty_stack", func(t *testing.T) {
		r := require.New(t)

		const defaultValue = 0

		ns := NewNodeStack[int]()

		empty := ns.IsEmpty()
		r.True(empty)

		val, ok := ns.Pop()
		r.False(ok)
		r.Equal(defaultValue, val)

		val, ok = ns.Pop()
		r.False(ok)
		r.Equal(defaultValue, val)
	})

	t.Run("popping_from_non-empty_stack", func(t *testing.T) {
		r := require.New(t)

		const defaultValue = 0

		ns := NewNodeStack[int](1, 2, 3)

		empty := ns.IsEmpty()
		r.False(empty)

		val, ok := ns.Pop()
		r.True(ok)
		r.Equal(3, val)

		val, ok = ns.Pop()
		r.True(ok)
		r.Equal(2, val)

		ns.Push(777)
		val, ok = ns.Top()
		r.True(ok)
		r.Equal(777, val)

		val, ok = ns.Pop()
		r.True(ok)
		r.Equal(777, val)

		val, ok = ns.Pop()
		r.True(ok)
		r.Equal(1, val)

		val, ok = ns.Pop()
		r.False(ok)
		r.Equal(defaultValue, val)

		empty = ns.IsEmpty()
		r.True(empty)
	})
}
