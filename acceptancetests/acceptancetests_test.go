package acceptancetests

import (
	"net/http"
	"testing"
	"time"
)

const (
	port = "8080"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	timout := 3 * time.Second // because each request processing in 2 sec
	cleanup, sendInterrupt, err := LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	// just check the server works before we shut things down
	assertCanGet(t, url, timout)
	// fire off a request, and before it has a chance to respond send SIGTERM.
	time.AfterFunc(50*time.Millisecond, func() {
		assertNoError(t, sendInterrupt())
	})
	// Without graceful shutdown, this would fail
	assertCanGet(t, url, timout)

	// after interrupt, the server should be shutdown, and no more requests will work
	assertCanNotGet(t, url, timout)
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("want no error, but got: %v", err)
	}
}

func assertCanNotGet(t testing.TB, url string, timoutInSec time.Duration) {
	t.Helper()
	if canGet(url, timoutInSec) {
		t.Errorf("want can't get, but can get by url: %s", url)
	}
}

func assertCanGet(t testing.TB, url string, timout time.Duration) {
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
