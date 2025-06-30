package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        testing.TB
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store was canceled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-data:
		return result, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("not cancelled request should write response", func(t *testing.T) {
		want := "Hello World"
		store := &SpyStore{t: t, response: want}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		writer := httptest.NewRecorder()

		server.ServeHTTP(writer, request)
		got := writer.Body.String()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("cancelled request should not write response", func(t *testing.T) {
		want := "hello, world"
		store := &SpyStore{t: t, response: want}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		writer := &SpyResponseWriter{}

		server.ServeHTTP(writer, request)
		if writer.written {
			t.Errorf("got response writer should not write response")
		}
	})
}
