package acceptancetests

import (
	"github.com/gnori-zon/go-tdd/generics/assert"
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
	assert.CanGet(t, url, timout)
	// fire off a request, and before it has a chance to respond send SIGTERM.
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})
	// Without graceful shutdown, this would fail
	assert.CanGet(t, url, timout)

	// after interrupt, the server should be shutdown, and no more requests will work
	assert.CanNotGet(t, url, timout)
}
