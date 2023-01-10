package stack

import "golang.org/x/exp/constraints"

var _ Stack[int] = &NodeStack[int]{}

type node[T constraints.Ordered] struct {
	val  T
	next *node[T]
}

// NodeStack is an implementation of Stack based on linked-list.
type NodeStack[T constraints.Ordered] struct {
	first *node[T]
}

func NewNodeStack[T constraints.Ordered](data ...T) *NodeStack[T] {
	ns := &NodeStack[T]{}
	for _, el := range data {
		ns.Push(el)
	}
	return ns
}

func (ns *NodeStack[T]) Push(val T) {
	if ns == nil {
		return
	}
	newNode := &node[T]{val: val}
	newNode.next = ns.first
	ns.first = newNode
}

func (ns *NodeStack[T]) Pop() (T, bool) {
	if ns.IsEmpty() {
		return *new(T), false
	}
	val := ns.first.val
	ns.first = ns.first.next
	return val, true
}

func (ns *NodeStack[T]) Top() (T, bool) {
	if ns.IsEmpty() {
		return *new(T), false
	}
	return ns.first.val, true
}

func (ns *NodeStack[T]) IsEmpty() bool {
	return ns == nil || ns.first == nil
}
