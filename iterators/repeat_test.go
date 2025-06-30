package iterators

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat when value is 'a' should return 'aaaaa'(5)", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertRepeated(t, expected, repeated)
	})

	t.Run("Repeat when value is 'b' should return 'bbbbbb'(6)", func(t *testing.T) {
		repeated := Repeat("b", 6)
		expected := "bbbbbb"
		assertRepeated(t, expected, repeated)
	})
}

func ExampleRepeat() {
	repeated := Repeat("c", 5)
	fmt.Println(repeated)
	// output: ccccc
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertRepeated(t testing.TB, expected, repeated string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("Expected %q, but got %q", expected, repeated)
	}
}
