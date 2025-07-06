package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(reader io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(reader).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league: %v", err)
	}
	return league, err
}

func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}
	return nil
}
