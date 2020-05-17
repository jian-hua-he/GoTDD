package main

import (
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var walkCount []string

	x := struct {
		Name string
	}{expected}

	Walk(x, func(input string) {
		walkCount = append(walkCount, input)
	})

	want := 1
	got := len(walkCount)
	if got != want {
		t.Errorf("wrong number of function calls, got %d wnat %d", got, want)
	}

	if walkCount[0] != expected {
		t.Errorf("got %q, want %q", walkCount[0], expected)
	}
}
