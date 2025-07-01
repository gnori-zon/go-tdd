package assert

import (
	"net/http"
	"testing"
	"time"
)

func NotEqual[T comparable](t testing.TB, lhs, rhs T) {
	t.Helper()
	if lhs == rhs {
		t.Errorf("expected not equals but %+v equal %+v", lhs, rhs)
	}
}

func Equal[T comparable](t testing.TB, lhs, rhs T) {
	t.Helper()
	if lhs != rhs {
		t.Errorf("expected equals but %+v not equal %+v", lhs, rhs)
	}
}

func True(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func NoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("want no error, but got: %v", err)
	}
}

func CanNotGet(t testing.TB, url string, timoutInSec time.Duration) {
	t.Helper()
	if canGet(url, timoutInSec) {
		t.Errorf("want can't get, but can get by url: %s", url)
	}
}

func CanGet(t testing.TB, url string, timout time.Duration) {
	t.Helper()
	if !canGet(url, timout) {
		t.Errorf("want can get, but can't get by url: %s", url)
	}
}

func canGet(url string, timeout time.Duration) bool {
	errChan := make(chan error)

	go func() {
		res, err := http.Get(url)
		if err != nil {
			errChan <- err
			return
		}
		_ = res.Body.Close()
		errChan <- nil
	}()
	select {
	case err := <-errChan:
		return err == nil
	case <-time.After(timeout):
		return false
	}
}
