package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	for _, player := range f.GetLeague() {
		if player.Name == name {
			return player.Wins
		}
	}

	return 0
}

func (f *FileSystemPlayerStore) Records(name string) {
	league := f.GetLeague()
	for i, player := range league {
		if player.Name == name {
			league[i].Wins += 1
			break
		}
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
