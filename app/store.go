package main

type PlayerStore interface {
	GetPlayerScore(player string) (score int, ok bool)
	SaveWin(player string)
}
