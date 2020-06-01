package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
    store := StubPlayerStore{
        map[string]int{
            "Pepper": 20,
            "Floyd": 10,
        },
    }
    server := &PlayerServer{&store}

	t.Run("returns Pepper’s score", func(test *testing.T) {
		req := newGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

        server.ServeHTTP(resp, req)

		assertResponseBody(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd’s score", func(test *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

        server.ServeHTTP(resp, req)

		assertResponseBody(t, resp.Body.String(), "10")
	})

    t.Run("returns 404 on missing players", func(t *testing.T) {
        req := newGetScoreRequest("Apollo")
        resp := httptest.NewRecorder()

        server.ServeHTTP(resp, req)

        got := resp.Code
        want := http.StatusNotFound

        if got != want {
            t.Errorf("got status %d want %d", got, want)
        }
    })
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
