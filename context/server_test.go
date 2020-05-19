package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: "hello, world"}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		ctx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(ctx)

		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, req)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
}
