// Package stack implements stack data structure
package stack

// Stack type that can take any kind of data structure.
// uses go generics
type Stack[T any] struct {
	elements []T
}

// New creates a new stack data structure of any type.
// returns the pointer new stack
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push pushes an elemt into the stack
// doesn't return anything
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop pops the element at the last index of the stack
// returns the element and bool
func (s *Stack[T]) Pop() (T, bool) {

	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}
