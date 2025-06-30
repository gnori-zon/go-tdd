package di

import (
	"bytes"
	"testing"
)

func TestGreat(t *testing.T) {
	t.Run("great should print hello people", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := Great(&buffer, "Chris")
		want := "Hello, Chris"

		got := buffer.String()

		if err != nil {
			t.Errorf("want nil err but got %v", err)
		}
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
