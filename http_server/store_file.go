package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
