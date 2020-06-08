package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := League(f.GetLeague()).Find(name)
	if player == nil {
		return 0
	}

	return player.Wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player == nil {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	} else {
		player.Wins += 1
	}

	f.database.Encode(f.league)
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
