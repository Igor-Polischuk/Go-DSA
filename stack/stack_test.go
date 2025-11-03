package stack

import (
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	t.Run("Test size", func(t *testing.T) {
		if stack.Size() == 0 {
			t.Errorf("expected size 2 after Push(1) and Push(2)")
		}
	})

	t.Run("Test peek", func(t *testing.T) {
		if stack.Peek() == 0 {
			t.Errorf("expected peek to be 4, got %d", stack.Peek())
		}
	})

	pop1 := stack.Pop()
	pop2 := stack.Pop()
	pop3 := stack.Pop()

	t.Run("Test pop", func(t *testing.T) {
		if pop1 != 4 && pop2 != 3 && pop3 != 2 {
			t.Errorf("expected pop to be 4, 3, 2, got %d, %d, %d", pop1, pop2, pop3)
		}
	})

	t.Run("Test is empty", func(t *testing.T) {
		if stack.IsEmpty() {
			t.Errorf("expected stack to be not empty, got %d", stack.Size())
		}

		stack.Pop()
		if !stack.IsEmpty() {
			t.Errorf("expected stack to be empty, got %d", stack.Size())
		}
	})

}
