package poker

import "testing"

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
