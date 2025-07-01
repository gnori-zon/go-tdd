package specification

import (
	"github.com/gnori-zon/go-tdd/generics/assert"
	"testing"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greater Greeter) {
	got, err := greater.Greet("Mike")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Mike!")
}

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t *testing.T, meany MeanGreeter) {
	got, err := meany.Curse("Chris")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go to hell, Chris!")
}
