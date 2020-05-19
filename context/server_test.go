package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy stoe got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestHandler(t *testing.T) {
	// t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
	// 	data := "hello, world"
	// 	store := &SpyStore{response: data}
	// 	server := Server(store)

	// 	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// 	ctx, cancel := context.WithCancel(req.Context())
	// 	time.AfterFunc(5*time.Millisecond, cancel)
	// 	req = req.WithContext(ctx)

	// 	resp := httptest.NewRecorder()
	// 	server.ServeHTTP(resp, req)

	// 	if !store.cancelled {
	// 		t.Errorf("store was not told to cancel")
	// 	}
	// })

	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		if resp.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, resp.Body.String(), data)
		}
	})
}
