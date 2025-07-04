package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
}

func NewPlayerServer() *PlayerServer {
	return &PlayerServer{NewInMemoryPlayerStore()}
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	player := extractPlayerName(request.URL.Path)
	switch request.Method {
	case http.MethodGet:
		s.printScore(w, player)
	case http.MethodPost:
		s.saveWin(w, player)
	}
}

func (s *PlayerServer) printScore(w http.ResponseWriter, player string) {
	score, ok := s.store.GetPlayerScore(player)
	if ok {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, score)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s *PlayerServer) saveWin(w http.ResponseWriter, player string) {
	w.WriteHeader(http.StatusAccepted)
	s.store.SaveWin(player)
}

func extractPlayerName(url string) string {
	return strings.TrimPrefix(url, "/players/")
}
