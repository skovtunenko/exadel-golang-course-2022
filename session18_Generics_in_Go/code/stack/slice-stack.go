package stack

import "golang.org/x/exp/constraints"

var _ Stack[int] = &SliceStack[int]{}

var _ Stack[string] = &SliceStack[string]{}

// SliceStack is an implementation of Stack based on slices.
type SliceStack[T constraints.Ordered] struct {
	data []T
}

// NewSliceStack constructs new Stack based internally on slices.
func NewSliceStack[T constraints.Ordered](data ...T) *SliceStack[T] {
	return &SliceStack[T]{data: data}
}

func (ss *SliceStack[T]) IsEmpty() bool {
	if ss == nil {
		return true
	}
	return len(ss.data) == 0
}

func (ss *SliceStack[T]) Push(elem T) {
	ss.data = append(ss.data, elem)
}

func (ss *SliceStack[T]) Pop() (T, bool) {
	if len(ss.data) == 0 {
		return *new(T), false
	}
	res := ss.data[len(ss.data)-1]
	ss.data = ss.data[:len(ss.data)-1]
	return res, true
}

func (ss *SliceStack[T]) Top() (T, bool) {
	if len(ss.data) == 0 {
		return *new(T), false
	}
	return ss.data[len(ss.data)-1], true
}
