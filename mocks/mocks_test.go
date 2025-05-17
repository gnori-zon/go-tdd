package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

type StubSleeper struct {
}

func (s StubSleeper) Sleep() {
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountDown(t *testing.T) {
	t.Run("should write correct message", func(t *testing.T) {
		buffer := bytes.Buffer{}
		CountDown(&buffer, &StubSleeper{})
		want := "3\n2\n1\nGO!"

		got := buffer.String()
		assertPrintString(t, want, got)
	})

	t.Run("should sleep before every next write after first write ", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	wantSleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{wantSleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != wantSleepTime {
		t.Errorf("should have slept for %v but slept for %v", wantSleepTime, spyTime.durationSlept)
	}
}

func assertPrintString(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want print %q but got %q", want, got)
	}
}
