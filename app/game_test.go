package poker

import (
	"fmt"
	"io"
	"testing"
	"time"
)

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &StubPlayerStore{}

func TestGame_Start(t *testing.T) {
	t.Run("it schedules printing of blind values for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := NewGame(dummyPlayerStore, blindAlerter)

		game.Start(5, io.Discard)

		wantAlerts := []Alert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		AssertPresentBlindAlerts(t, wantAlerts, blindAlerter.Alerts)
	})

	t.Run("it schedules printing of blind values for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := NewGame(dummyPlayerStore, blindAlerter)

		game.Start(7, io.Discard)

		wantAlerts := []Alert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		AssertPresentBlindAlerts(t, wantAlerts, blindAlerter.Alerts)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &StubPlayerStore{}
	game := NewGame(store, dummyBlindAlerter)
	winner := "Ruth"

	game.Finish(winner)
	AssertSavedWin(t, store.Wins, 1, winner)
}

func AssertPresentBlindAlerts(t *testing.T, wantAlerts, gotAlerts []Alert) {
	t.Helper()
	for i, want := range wantAlerts {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(gotAlerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, gotAlerts)
			}

			got := gotAlerts[i]
			AssertEqualAlert(t, got, want)
		})
	}
}
