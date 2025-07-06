package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"strconv"
	"sync"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	wins   []string
	league []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) (score int, ok bool) {
	score, ok = s.scores[name]
	return
}

func (s *StubPlayerStore) SaveWin(name string) {
	s.wins = append(s.wins, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	server := NewPlayerServer(&StubPlayerStore{scores: map[string]int{"Pepper": 20, "Bob": 10}, wins: []string{}})
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
	stubStore := &StubPlayerStore{scores: map[string]int{}, wins: []string{}}
	server := NewPlayerServer(stubStore)

	wantCountWins := 0

	for _, name := range []string{"Pepper", "Bob"} {
		wantCountWins++
		t.Run(fmt.Sprintf("save win for %s", name), func(t *testing.T) {
			request := newSaveWinRequest(name)
			pepperResponse := httptest.NewRecorder()

			server.ServeHTTP(pepperResponse, request)

			assertStatus(t, pepperResponse, http.StatusAccepted)
			assertSavedWin(t, stubStore.wins, wantCountWins, name)
		})
	}
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	league := []Player{
		{Name: "Pepper", Wins: 12},
		{Name: "Bob", Wins: 40},
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
	server := NewPlayerServer(NewInMemoryPlayerStore())
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

		server := NewPlayerServer(&StubPlayerStore{scores: map[string]int{}, wins: []string{}, league: wantLeague})

		request := newGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertEqualLeagues(t, response, wantLeague)
		assertContentType(t, response, jsonContentType)
		assertStatus(t, response, http.StatusOK)
	})
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

func newSaveWinRequest(bobName string) *http.Request {
	req, _ := http.NewRequest("POST", fmt.Sprintf("/players/%s", bobName), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertSavedWin(t testing.TB, wins []string, wantCountWins int, name string) {
	t.Helper()
	if len(wins) != wantCountWins || wins[wantCountWins-1] != name {
		t.Errorf("store wins for %q not stored", name)
	}
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
