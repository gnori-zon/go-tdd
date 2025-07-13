package cli

import (
	poker "github.com/gnori-zon/go-tdd/app"
	"time"
)

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type TexasHoldem struct {
	store        poker.PlayerStore
	blindAlerter BlindAlerter
}

func NewGame(store poker.PlayerStore, blindAlerter BlindAlerter) *TexasHoldem {
	return &TexasHoldem{store, blindAlerter}
}

func (g TexasHoldem) Start(numberOfPlayers int) {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Minute
	timeToIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	for _, blind := range blinds {
		g.blindAlerter.ScheduleAlertAt(blindTime, blind)
		blindTime += timeToIncrement
	}
}

func (g TexasHoldem) Finish(winner string) {
	g.store.SaveWin(winner)
}
