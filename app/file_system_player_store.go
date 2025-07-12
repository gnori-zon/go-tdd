package poker

import (
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"sync"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
	mu       sync.Mutex
}

func NewFileSystemPlayerStoreFromFile(databasePath string) (store *FileSystemPlayerStore, close func(), err error) {
	database, err := os.OpenFile(databasePath, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", databasePath, err)
	}
	close = func() {
		_ = database.Close()
	}
	store, err = NewFileSystemPlayerStore(database)
	return
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	err := prepareDatabase(database)
	if err != nil {
		return nil, fmt.Errorf("could not prepare database for file system player store: %w", err)
	}
	league, err := NewLeague(database)
	if err != nil {
		return nil, fmt.Errorf("could not create league: %w", err)
	}
	return &FileSystemPlayerStore{database: json.NewEncoder(&tape{database}), league: league}, nil
}

func prepareDatabase(database *os.File) error {
	_, _ = database.Seek(0, io.SeekStart)
	info, err := database.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat db file: %w", err)
	}
	if info.Size() == 0 {
		_, _ = database.Write([]byte("[]"))
		_, _ = database.Seek(0, io.SeekStart)
	}
	return nil
}

func (s *FileSystemPlayerStore) GetLeague() League {
	slices.SortFunc(s.league, func(a, b Player) int {
		return cmp.Compare(b.Wins, a.Wins)
	})
	return s.league
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) (int, bool) {
	player := s.league.Find(name)
	if player != nil {
		return player.Wins, true
	}
	return -1, false
}

func (s *FileSystemPlayerStore) SaveWin(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	player := s.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		s.league = append(s.league, Player{name, 1})
	}
	_ = s.database.Encode(s.league)
}
