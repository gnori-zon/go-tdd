package poker

import (
	"fmt"
	"io"
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

func (s *SpyBlindAlerter) ScheduleAlertAt(scheduledAt time.Duration, amount int, alertDestination io.Writer) {
	s.Alerts = append(s.Alerts, Alert{scheduledAt, amount})
}

type SpyGame struct {
	StartCalled  bool
	FinishCalled bool
	BlindAlert   []byte
	StartWith    int
	FinishWith   string
}

func (s *SpyGame) Start(numberOfPlayers int, alertDestination io.Writer) {
	s.StartCalled = true
	s.StartWith = numberOfPlayers
	_, _ = alertDestination.Write(s.BlindAlert)
}

func (s *SpyGame) Finish(winner string) {
	s.FinishCalled = true
	s.FinishWith = winner
}

func AssertGameStartWith(t *testing.T, game *SpyGame, wantNumberOfPlayers int) {
	if game.StartWith != wantNumberOfPlayers {
		t.Errorf("want start with %d, but got %d", wantNumberOfPlayers, game.StartWith)
	}
}

func AssertGameFinishWith(t testing.TB, game *SpyGame, wantName string) {
	t.Helper()
	if game.FinishWith != wantName {
		t.Errorf("want finish game with %q, but got %q", wantName, game.FinishWith)
	}
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
