package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response  string
	cancelled bool
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "hello, world"
	server := Server(&StubStore{data, false})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	if resp.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, resp.Body.String(), data)
	}
}
