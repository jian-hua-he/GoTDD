package main

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := map[string]struct {
		Arabic int
		Want   string
	}{
		"1 gets converted to I": {
			Arabic: 1,
			Want:   "I",
		},
		"2 gets converted to II": {
			Arabic: 2,
			Want:   "II",
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {

			got := ConvertToRoman(c.Arabic)
			want := c.Want

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
