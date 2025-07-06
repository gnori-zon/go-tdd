package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

const jsonContentType = "application/json"

func NewPlayerServer(store PlayerStore) *PlayerServer {
	server := new(PlayerServer)
	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(server.playersHandler()))
	router.Handle("/league", http.HandlerFunc(server.leagueHandler()))

	server.Handler = router
	server.store = store
	return server
}

// region leagueHandler
func (s *PlayerServer) leagueHandler() func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		_ = json.NewEncoder(w).Encode(s.store.GetLeague())
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", jsonContentType)
	}
}

// endregion

// region playersHandler
func (s *PlayerServer) playersHandler() func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		player := extractPlayerName(request.URL.Path)
		switch request.Method {
		case http.MethodGet:
			s.printScore(w, player)
		case http.MethodPost:
			s.saveWin(w, player)
		}
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

// endregion

type Player struct {
	Name string `json:"name"`
	Wins int    `json:"wins"`
}
