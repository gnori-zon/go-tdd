package intersection

import (
	"github.com/gnori-zon/go-tdd/scalingacceptancetests/specification"
	"testing"
)

func TestGreet(t *testing.T) {
	specification.GreetSpecification(t, specification.GreetAdapter(Greet))
}

func TestCurse(t *testing.T) {
	specification.CurseSpecification(t, specification.MeanGreetAdapter(Curse))
}
