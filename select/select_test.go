package _select

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("return faster url", func(t *testing.T) {
		slowServer := makeDelayedServer(5 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)
		if want != got {
			t.Errorf("want %q but got %q", want, got)
		}
	})

	t.Run("return error when best url slower than timeout", func(t *testing.T) {
		timeout := 2 * time.Millisecond
		slowServer := makeDelayedServer(3 * time.Millisecond)
		fastServer := makeDelayedServer(3 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		_, err := ConfigurableRacer(slowUrl, fastUrl, timeout)
		var gotErr ErrTimeout
		if !errors.As(err, &gotErr) {
			t.Errorf("want err timeout but got %q", err)
		}
	})
}

func BenchmarkRacer(b *testing.B) {
	slowServer := makeDelayedServer(2 * time.Millisecond)
	fastServer := makeDelayedServer(5 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Racer(slowUrl, fastUrl)
	}
}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
