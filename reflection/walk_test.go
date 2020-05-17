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
		"struct with one string field": {
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		"struct with two string field": {
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		"struct with non string field": {
			struct {
				Name string
				Age  int
			}{"Chris", 33},
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
