package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("get fastet url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but get one, err %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn’t respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(6 * time.Second)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 5*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn’t get one")
		}
	})
}

func makeDelayedServer(deplay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(deplay)
		w.WriteHeader(http.StatusOK)
	}))
}
