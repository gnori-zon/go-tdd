package specification

type GreetAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}

type MeanGreetAdapter func(name string) string

func (g MeanGreetAdapter) Curse(name string) (string, error) {
	return g(name), nil
}
