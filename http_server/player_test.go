package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

const jsonContentType = "application/json"

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		make([]string, 0),
		make([]Player, 0),
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper’s score", func(test *testing.T) {
		req := newGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd’s score", func(test *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("Apollo")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusNotFound)
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func TestStoreWins(t *testing.T) {
	t.Run("it returns accepted on POST", func(t *testing.T) {
		store := StubPlayerStore{
			map[string]int{},
			make([]string, 0),
			make([]Player, 0),
		}
		server := NewPlayerServer(&store)

		req := newPostWinRequest("Pepper")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusAccepted)
	})

	t.Run("it records wins on POST", func(t *testing.T) {
		store := StubPlayerStore{
			map[string]int{},
			make([]string, 0),
			make([]Player, 0),
		}
		server := NewPlayerServer(&store)

		player := "Pepper"
		req := newPostWinRequest(player)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
	return req
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, newGetScoreRequest(player))
		assertStatus(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, newLeagueRequest())
		assertStatus(t, resp.Code, http.StatusOK)

		got := getLeagueFromResponse(t, resp.Body)
		want := []Player{
			{"Pepper", 3},
		}

		assertLeague(t, got, want)
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		var got []Player
		if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", resp.Body, err)
		}

		assertStatus(t, resp.Code, http.StatusOK)
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		want := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, want}
		server := NewPlayerServer(&store)

		req := newLeagueRequest()
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := getLeagueFromResponse(t, resp.Body)
		assertStatus(t, resp.Code, http.StatusOK)
		assertLeague(t, got, want)
		assertContentType(t, resp, jsonContentType)
	})
}

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}
        ]`)
		defer cleanDb()

		store := FileSystemPlayerStore{db}
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// Run again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}
        ]`)
		defer cleanDb()

		store := FileSystemPlayerStore{db}

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store win for existing players", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}
        ]`)
		defer cleanDb()

		store := FileSystemPlayerStore{db}
		store.Records("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEquals(t, got, want)
	})

	t.Run("store win for new players", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
	        {"Name": "Chris", "Wins": 33}
	    ]`)
		defer cleanDb()

		store := FileSystemPlayerStore{db}
		store.Records("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tempFile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempFile.Write([]byte(initialData))

	removeFile := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("unable to parse response from server %q into slice of Player, %v", body, err)
	}

	return
}

func assertContentType(t *testing.T, resp *httptest.ResponseRecorder, want string) {
	t.Helper()
	if resp.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, resp.Result().Header)
	}
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
