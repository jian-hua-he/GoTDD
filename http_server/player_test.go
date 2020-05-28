package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper’s score", func(test *testing.T) {
		req := getNewScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)

		assertResponseBody(resp.Body.String(), "20")
	})

	t.Run("returns Floyd’s score", func(test *testing.T) {
		req := getNewScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)

		assertResponseBody(resp.Body.String(), "10")
	})
}

func getNewScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
