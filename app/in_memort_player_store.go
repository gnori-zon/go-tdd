package main

import "sync"

type InMemoryPlayerStore struct {
	mu    sync.Mutex
	score map[string]int
}

func NewInMemoryPlayerStore() PlayerStore {
	return &InMemoryPlayerStore{score: make(map[string]int)}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) (score int, ok bool) {
	score, ok = s.score[name]
	return
}

func (s *InMemoryPlayerStore) SaveWin(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.score[name]++
}
