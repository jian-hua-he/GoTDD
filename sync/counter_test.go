package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	counter := Counter{}
	counter.Inc()
	counter.Inc()
	counter.Inc()

	if counter.Value() != 3 {
		t.Errorf("got %v, want %v", counter.Value(), 3)
	}
}
