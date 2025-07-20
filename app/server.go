package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

type PlayerServer struct {
	store   PlayerStore
	game    Game
	gameTmp *template.Template
	http.Handler
}

const jsonContentType = "application/json"
const htmlContentType = "text/html"

func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	gameTmp, err := template.ParseFiles("game.html")
	if err != nil {
		err = fmt.Errorf("problem loading template %s", err.Error())
		return nil, err
	}

	server := new(PlayerServer)
	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(server.playersHandler()))
	router.Handle("/league", http.HandlerFunc(server.leagueHandler()))
	router.Handle("/game", http.HandlerFunc(server.gameHandler()))
	router.Handle("/ws", http.HandlerFunc(server.wsHandler()))

	server.Handler = router
	server.game = game
	server.gameTmp = gameTmp
	server.store = store
	return server, nil
}

func mustMakePlayerServer(t testing.TB, store PlayerStore, game Game) *PlayerServer {
	server, err := NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
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

// region gameHandler
func (s *PlayerServer) gameHandler() func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", htmlContentType)
		_ = s.gameTmp.Execute(w, nil)
	}
}

// endregion

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *PlayerServer) wsHandler() func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		ws := newPlayerServerWS(w, request)
		numberOfPlayersMsg := ws.WaitForMsg()
		numberOfPlayers, _ := strconv.Atoi(numberOfPlayersMsg)
		s.game.Start(numberOfPlayers, ws)

		winner := ws.WaitForMsg()
		s.game.Finish(winner)
	}
}

func extractPlayerName(url string) string {
	return strings.TrimPrefix(url, "/players/")
}

// endregion

// region playerServerWs
type playerServerWs struct {
	*websocket.Conn
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWs {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("problem upgrading connection to WebSockets %v\n", err)
	}

	return &playerServerWs{conn}
}

func (w *playerServerWs) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket %v\n", err)
	}
	return string(msg)
}
func (w *playerServerWs) Write(payload []byte) (n int, err error) {
	err = w.WriteMessage(websocket.TextMessage, payload)
	if err != nil {
		return 0, err
	}
	return len(payload), nil
}

// endregion

type Player struct {
	Name string `json:"name"`
	Wins int    `json:"wins"`
}
