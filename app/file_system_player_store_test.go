package poker

import (
	"encoding/json"
	"github.com/gnori-zon/go-tdd/generics/assert"
	"os"
	"slices"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("get sorted league", func(t *testing.T) {
		chris := Player{Name: "Chris", Wins: 33}
		cleo := Player{Name: "Cleo", Wins: 10}
		database, cleanDatabase := createTempFile(t, []Player{cleo, chris})
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(database)

		want := []Player{chris, cleo}
		got := store.GetLeague()
		assertLeague(t, want, got)

		got = store.GetLeague()
		assertLeague(t, want, got)
	})

	t.Run("get player score", func(t *testing.T) {
		name := "Cleo"
		want := 10
		database, cleanDatabase := createTempFile(t, []Player{{Name: name, Wins: want}, {Name: "Chris", Wins: 33}})
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(database)

		got, ok := store.GetPlayerScore(name)
		assertFoundedPlayerScore(t, name, ok, got, want)

		_, ok = store.GetPlayerScore("not exist name")
		assertNotFoundedPlayerScore(t, name, ok)
	})

	t.Run("store player score for existing players", func(t *testing.T) {
		name := "Cleo"
		oldWins := 10
		database, cleanDatabase := createTempFile(t, []Player{{Name: name, Wins: oldWins}, {Name: "Chris", Wins: 33}})
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(database)

		store.SaveWin(name)
		got, _ := store.GetPlayerScore(name)

		want := oldWins + 1
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("store player score for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, []Player{})
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(database)

		name := "Bob"
		store.SaveWin(name)
		got, _ := store.GetPlayerScore(name)

		want := 1
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFileFromString(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		assert.NoError(t, err)
	})
}

func assertNotFoundedPlayerScore(t testing.TB, name string, ok bool) {
	t.Helper()
	if ok {
		t.Errorf("want not found score for player: %s, but is founded", name)
	}
}

func assertFoundedPlayerScore(t *testing.T, name string, ok bool, got int, want int) {
	t.Helper()
	if !ok {
		t.Errorf("want found score for player: %s, but got nothing", name)
	}
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertLeague(t *testing.T, want []Player, got []Player) {
	t.Helper()
	if !slices.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func createTempFileFromString(t testing.TB, content string) (tempFile *os.File, removeFile func()) {
	t.Helper()
	tempFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatal(err)
	}
	removeFile = func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	_, err = tempFile.Write([]byte(content))
	if err != nil {
		removeFile()
		t.Fatal(err)
	}
	return tempFile, removeFile
}

func createTempFile(t testing.TB, players []Player) (tempFile *os.File, removeFile func()) {
	t.Helper()
	tempFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatal(err)
	}
	removeFile = func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	err = json.NewEncoder(tempFile).Encode(players)
	if err != nil {
		removeFile()
		t.Fatal(err)
	}
	return tempFile, removeFile
}
