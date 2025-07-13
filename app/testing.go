package poker

import (
	"fmt"
	"testing"
	"time"
)

type StubPlayerStore struct {
	Scores map[string]int
	Wins   []string
	League []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) (score int, ok bool) {
	score, ok = s.Scores[name]
	return
}

func (s *StubPlayerStore) SaveWin(name string) {
	s.Wins = append(s.Wins, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertSavedWin(t testing.TB, wins []string, wantCountWins int, name string) {
	t.Helper()
	if len(wins) != wantCountWins {
		t.Errorf("store wins for %q not stored", name)
	}
	got := wins[wantCountWins-1]
	if got != name {
		t.Errorf("wanted win for %q but got %q", name, got)
	}
}

type Alert struct {
	ScheduledAt time.Duration
	Amount      int
}

func (a Alert) String() string {
	return fmt.Sprintf("%d scheduled for %v", a.Amount, a.ScheduledAt)
}

type SpyBlindAlerter struct {
	Alerts []Alert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(scheduledAt time.Duration, amount int) {
	s.Alerts = append(s.Alerts, Alert{scheduledAt, amount})
}

func AssertEqualAlert(t testing.TB, got Alert, want Alert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf("want alert amount %d but got %d", want.Amount, got.Amount)
	}
	if got.ScheduledAt != want.ScheduledAt {
		t.Errorf("want alert scheduled time %v but got %v", want.ScheduledAt, got.ScheduledAt)
	}
}
