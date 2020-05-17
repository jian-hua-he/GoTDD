package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := map[string]struct {
		Input         interface{}
		ExpectedCalls []string
	}{
		"Struct with one string field": {
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			var got []string
			Walk(c.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, c.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, c.ExpectedCalls)
			}
		})
	}
}
