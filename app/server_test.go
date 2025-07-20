package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httptest"
	"slices"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	dummyGame = &SpyGame{}
)

func TestGETPlayers(t *testing.T) {
	server := mustMakePlayerServer(t, &StubPlayerStore{Scores: map[string]int{"Pepper": 20, "Bob": 10}, Wins: []string{}}, dummyGame)
	cases := []struct {
		name      string
		wantScore int
	}{
		{
			name:      "Pepper",
			wantScore: 20,
		},
		{
			name:      "Bob",
			wantScore: 10,
		},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("return %s's score for exist players", testCase.name), func(t *testing.T) {
			request := newGetScoreRequest(testCase.name)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)
			assertScore(t, response, testCase.wantScore)
			assertStatus(t, response, http.StatusOK)
		})
	}

	t.Run("return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("unkown1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	stubStore := &StubPlayerStore{Scores: map[string]int{}, Wins: []string{}}
	server := mustMakePlayerServer(t, stubStore, dummyGame)

	wantCountWins := 0

	for _, name := range []string{"Pepper", "Bob"} {
		wantCountWins++
		t.Run(fmt.Sprintf("save win for %s", name), func(t *testing.T) {
			request := newSaveWinRequest(name)
			pepperResponse := httptest.NewRecorder()

			server.ServeHTTP(pepperResponse, request)

			assertStatus(t, pepperResponse, http.StatusAccepted)
			AssertSavedWin(t, stubStore.Wins, wantCountWins, name)
		})
	}
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, []Player{})
	defer cleanDatabase()
	store, _ := NewFileSystemPlayerStore(database)
	server := mustMakePlayerServer(t, store, dummyGame)
	league := []Player{
		{Name: "Bob", Wins: 40},
		{Name: "Pepper", Wins: 12},
	}

	for _, player := range league {
		for i := 0; i < player.Wins; i++ {
			server.ServeHTTP(httptest.NewRecorder(), newSaveWinRequest(player.Name))
		}

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player.Name))
		assertStatus(t, response, http.StatusOK)
		assertScore(t, response, player.Wins)
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetLeagueRequest())

	assertEqualLeagues(t, response, league)
	assertStatus(t, response, http.StatusOK)
	assertContentType(t, response, jsonContentType)
}

func TestConcurrentlyRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, []Player{})
	defer cleanDatabase()
	store, _ := NewFileSystemPlayerStore(database)
	server := mustMakePlayerServer(t, store, dummyGame)
	player := "Pepper"

	want := 1_000
	var wait sync.WaitGroup
	wait.Add(want)
	for i := 0; i < want; i++ {
		go func() {
			server.ServeHTTP(httptest.NewRecorder(), newSaveWinRequest(player))
			wait.Done()
		}()
	}
	wait.Wait()

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response, http.StatusOK)
	assertScore(t, response, want)
}

func TestLeague(t *testing.T) {
	t.Run("return league list", func(t *testing.T) {
		wantLeague := []Player{
			{Name: "Pepper", Wins: 12},
			{Name: "Bob", Wins: 20},
		}

		server := mustMakePlayerServer(t, &StubPlayerStore{Scores: map[string]int{}, Wins: []string{}, League: wantLeague}, dummyGame)

		request := newGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertEqualLeagues(t, response, wantLeague)
		assertContentType(t, response, jsonContentType)
		assertStatus(t, response, http.StatusOK)
	})
}

const tenMs = 10 * time.Millisecond

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &StubPlayerStore{Scores: map[string]int{}, Wins: []string{}}, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
		assertContentType(t, response, htmlContentType)
		body := response.Body.String()
		if !strings.Contains(body, "Lets play poker") {
			t.Errorf("response body does not contain html: %s", body)
		}
	})

	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		wantBlindAlert := "Blind is 100"
		playerStore := &StubPlayerStore{}
		game := &SpyGame{BlindAlert: []byte(wantBlindAlert)}
		winner := "Ruth"
		wantCountPlayers := 3
		server := httptest.NewServer(mustMakePlayerServer(t, playerStore, game))
		ws := mustDialWs(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")
		defer server.Close()
		defer ws.Close()

		writeWsMessage(t, ws, strconv.Itoa(wantCountPlayers))
		writeWsMessage(t, ws, winner)

		time.Sleep(tenMs)
		retryAssertGameStartWith(t, game, wantCountPlayers)
		retryAssertGameFinishWith(t, game, winner)
		within(t, tenMs, func() {
			assertWsGotMessage(t, ws, wantBlindAlert)
		})
	})
}

func retryAssertGameStartWith(t testing.TB, game *SpyGame, wantNumberOfPlayers int) {
	t.Helper()
	passed := retryUntil(tenMs, func() bool {
		return game.StartWith == wantNumberOfPlayers
	})
	if !passed {
		t.Errorf("want start with %d, but got %d", wantNumberOfPlayers, game.StartWith)
	}
}

func retryAssertGameFinishWith(t testing.TB, game *SpyGame, wantWinner string) {
	t.Helper()
	passed := retryUntil(tenMs, func() bool {
		return game.FinishWith == wantWinner
	})
	if !passed {
		t.Errorf("want finish game with %q, but got %q", wantWinner, game.FinishWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}

func assertWsGotMessage(t testing.TB, ws *websocket.Conn, want string) {
	t.Helper()
	_, gotMsg, _ := ws.ReadMessage()
	got := string(gotMsg)
	if got != want {
		t.Errorf("want got ws msg %q, but got %q", want, got)
	}
}

func within(t testing.TB, duration time.Duration, assert func()) {
	t.Helper()
	done := make(chan struct{})

	go func() {
		assert()
		close(done)
	}()

	select {
	case <-time.After(duration):
		t.Errorf("timed out waiting for %s", duration)
	case <-done:
	}
}

func writeWsMessage(t testing.TB, ws *websocket.Conn, message string) {
	t.Helper()
	if err := ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send ws message: %s", err)
	}
}

func mustDialWs(t testing.TB, wsURL string) *websocket.Conn {
	t.Helper()
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("could not open ws connection: %s", err)
	}
	return ws
}

func assertEqualLeagues(t *testing.T, response *httptest.ResponseRecorder, want []Player) {
	t.Helper()
	got := getLeagueFromResponse(t, response)
	if !slices.Equal(want, got) {
		t.Errorf("want league %v, got %v", want, got)
	}
}

func newGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/league", nil)
	return req
}

func getLeagueFromResponse(t *testing.T, response *httptest.ResponseRecorder) []Player {
	t.Helper()
	var got []Player
	err := json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Fatalf("error decoding league response: %v", err)
	}
	return got
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	contentType := response.Header().Get("Content-Type")
	if contentType != want {
		t.Errorf("Content-Type is %q, want %q", contentType, want)
	}
}

func newGameRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/game", nil)
	return req
}

func newSaveWinRequest(bobName string) *http.Request {
	req, _ := http.NewRequest("POST", fmt.Sprintf("/players/%s", bobName), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertScore(t testing.TB, response *httptest.ResponseRecorder, wantScore int) {
	t.Helper()
	got := response.Body.String()
	want := strconv.Itoa(wantScore)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, response *httptest.ResponseRecorder, wantCode int) {
	t.Helper()
	if response.Code != wantCode {
		t.Errorf("got %d, want %d", response.Code, wantCode)
	}
}
