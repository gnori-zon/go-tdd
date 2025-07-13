package cli

import (
	"fmt"
	poker "github.com/gnori-zon/go-tdd/app"
	cli2 "github.com/gnori-zon/go-tdd/app/cli"
	"testing"
	"time"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

func TestGame_Start(t *testing.T) {
	t.Run("it schedules printing of blind values for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := cli2.NewGame(dummyPlayerStore, blindAlerter)

		game.Start(5)

		wantAlerts := []poker.Alert{
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
		blindAlerter := &poker.SpyBlindAlerter{}
		game := cli2.NewGame(dummyPlayerStore, blindAlerter)

		game.Start(7)

		wantAlerts := []poker.Alert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		AssertPresentBlindAlerts(t, wantAlerts, blindAlerter.Alerts)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := cli2.NewGame(store, dummyBlindAlerter)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertSavedWin(t, store.Wins, 1, winner)
}

func AssertPresentBlindAlerts(t *testing.T, wantAlerts, gotAlerts []poker.Alert) {
	t.Helper()
	for i, want := range wantAlerts {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(gotAlerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, gotAlerts)
			}

			got := gotAlerts[i]
			poker.AssertEqualAlert(t, got, want)
		})
	}
}
