package stack

import "golang.org/x/exp/constraints"

// Stack is a generic stack contract.
type Stack[T constraints.Ordered] interface {
	// Push adds value onto the stack.
	Push(val T)
	// Pop populates value from the stack (and removes it).
	Pop() (T, bool)
	// Top gets the topmost value without removing it from the Stack.
	Top() (T, bool)
	// IsEmpty checks if the stack is empty or not.
	IsEmpty() bool
}
