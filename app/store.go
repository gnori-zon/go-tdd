package poker

type PlayerStore interface {
	GetPlayerScore(player string) (score int, ok bool)
	SaveWin(player string)
	GetLeague() League
}
