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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChan := make(chan Profile)

		go func() {
			aChan <- Profile{33, "Berlin"}
			aChan <- Profile{23, "Katowice"}
			close(aChan)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aChan, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn’t", haystack, needle)
	}
}
