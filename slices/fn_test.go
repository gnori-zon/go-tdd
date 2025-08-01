package slices

import (
	"github.com/gnori-zon/go-tdd/generics/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		assert.Equal(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		assert.Equal(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}
