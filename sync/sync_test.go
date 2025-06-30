package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertValue(t, counter, 3)
	})

	t.Run("it run safely concurrently", func(t *testing.T) {
		wantCount := 1000
		counter := NewCounter()

		var waitGroup sync.WaitGroup
		waitGroup.Add(wantCount)
		for i := 0; i < wantCount; i++ {
			go func() {
				defer waitGroup.Done()
				counter.Inc()
			}()
		}
		waitGroup.Wait()

		assertValue(t, counter, wantCount)
	})
}

func assertValue(t testing.TB, counter *Counter, want int) {
	t.Helper()
	got := counter.Value()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
