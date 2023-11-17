package stack

import (
	"testing"
)

func TestPop(t *testing.T) {
	stack := New[int]()

	stack.Push(10)
	stack.Push(20)

	if val, ok := stack.Pop(); !ok || val != 20 {
		t.Errorf("Pop() != %d; want 20", val)
	}

	if val, ok := stack.Pop(); !ok || val != 10 {
		t.Errorf("Pop() = %d; want 10", val)
	}

	if _, ok := stack.Pop(); ok {
		t.Error("Pop() should be false when popping from empty stack")
	}
}
