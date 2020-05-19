package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestHandler(t *testing.T) {
	data := "hello, world"
	server := Server(&StubStore{data})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	if resp.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, resp.Body.String(), data)
	}
}
