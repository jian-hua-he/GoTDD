package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %v, want %v", counter.Value(), 3)
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i += 1 {
			go func(wg *sync.WaitGroup, counter *Counter) {
				counter.Inc()
				wg.Done()
			}(&wg, counter)
		}
		wg.Wait()

		if counter.Value() != wantedCount {
			t.Errorf("got %v, want %v", counter.Value(), wantedCount)
		}
	})
}
