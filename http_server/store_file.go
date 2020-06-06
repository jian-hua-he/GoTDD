package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)
	if player == nil {
		return 0
	}

	return player.Wins
}

func (f *FileSystemPlayerStore) Records(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins += 1
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
