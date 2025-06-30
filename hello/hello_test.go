package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Hello when name is not blank should return hello, people", func(t *testing.T) {
		language := ""
		got := Hello("Chris", language)
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want, language)
	})

	t.Run("Hello when name is blank should return hello, World", func(t *testing.T) {
		for _, emptyName := range []string{"", " "} {
			language := ""
			got := Hello(emptyName, language)
			want := "Hello, World"
			assertCorrectMessage(t, got, want, language)
		}
	})

	t.Run("Hello when language is not default should return Hola, people", func(t *testing.T) {
		languageWithPrefix := map[string]string{
			"Spanish": "Hola",
			"French":  "Bonjour",
		}
		for language, prefix := range languageWithPrefix {
			got := Hello("Elodie", language)
			want := fmt.Sprintf("%s, Elodie", prefix)
			assertCorrectMessage(t, got, want, language)
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want, language string) {
	t.Helper()
	if language == "" {
		language = "default"
	}
	if got != want {
		t.Errorf("got %q want %q, for language %q", got, want, language)
	}
}
