package generics

import (
	"github.com/gnori-zon/go-tdd/generics/assert"
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		assert.Equal(t, 1, 1)
		assert.NotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		assert.Equal(t, "1", "1")
		assert.NotEqual(t, "1", "2")
	})
}

func TestStack(t *testing.T) {

	t.Run("push should add new integer to stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)

		assert.Equal(t, 1, stack.Len())
	})

	t.Run("pop should remove integer from stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		removed, ok := stack.Pop()

		assert.True(t, ok)
		assert.Equal(t, 1, removed)
		assert.Equal(t, 0, stack.Len())
	})

	t.Run("push should add new string to stack", func(t *testing.T) {
		stack := NewStack[string]()
		stack.Push("1")

		assert.Equal(t, 1, stack.Len())
	})

	t.Run("pop should remove integer from stack", func(t *testing.T) {
		stack := NewStack[string]()
		stack.Push("1")
		removed, ok := stack.Pop()

		assert.True(t, ok)
		assert.Equal(t, "1", removed)
		assert.Equal(t, 0, stack.Len())
	})

	t.Run("push should add new element to head stack and removed from head stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		stack.Push(2)
		firstRemoved, ok := stack.Pop()

		assert.True(t, ok)
		assert.Equal(t, 2, firstRemoved)
		assert.Equal(t, 1, stack.Len())

		secondRemoved, ok := stack.Pop()

		assert.True(t, ok)
		assert.Equal(t, 1, secondRemoved)
		assert.Equal(t, 0, stack.Len())

		assert.Equal(t, firstRemoved+secondRemoved, 3)
	})
}
