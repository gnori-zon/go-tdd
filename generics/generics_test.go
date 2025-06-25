package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "1", "1")
		AssertNotEqual(t, "1", "2")
	})
}

func TestStack(t *testing.T) {

	t.Run("push should add new integer to stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)

		AssertEqual(t, 1, stack.Len())
	})

	t.Run("pop should remove integer from stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		removed, ok := stack.Pop()

		AssertTrue(t, ok)
		AssertEqual(t, 1, removed)
		AssertEqual(t, 0, stack.Len())
	})

	t.Run("push should add new string to stack", func(t *testing.T) {
		stack := NewStack[string]()
		stack.Push("1")

		AssertEqual(t, 1, stack.Len())
	})

	t.Run("pop should remove integer from stack", func(t *testing.T) {
		stack := NewStack[string]()
		stack.Push("1")
		removed, ok := stack.Pop()

		AssertTrue(t, ok)
		AssertEqual(t, "1", removed)
		AssertEqual(t, 0, stack.Len())
	})

	t.Run("push should add new element to head stack and removed from head stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		stack.Push(2)
		firstRemoved, ok := stack.Pop()

		AssertTrue(t, ok)
		AssertEqual(t, 2, firstRemoved)
		AssertEqual(t, 1, stack.Len())

		secondRemoved, ok := stack.Pop()

		AssertTrue(t, ok)
		AssertEqual(t, 1, secondRemoved)
		AssertEqual(t, 0, stack.Len())

		AssertEqual(t, firstRemoved+secondRemoved, 3)
	})
}

func AssertNotEqual[T comparable](t *testing.T, lhs, rhs T) {
	t.Helper()
	if lhs == rhs {
		t.Errorf("expected not equals but %+v equal %+v", lhs, rhs)
	}
}

func AssertEqual[T comparable](t *testing.T, lhs, rhs T) {
	t.Helper()
	if lhs != rhs {
		t.Errorf("expected equals but %+v not equal %+v", lhs, rhs)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
