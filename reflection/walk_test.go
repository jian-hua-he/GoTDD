package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		"nested fields": {
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		"pointers to things": {
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		"slices": {
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		"arrays": {
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
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
