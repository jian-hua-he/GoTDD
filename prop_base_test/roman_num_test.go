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
		"3 gets converted to III": {
			Arabic: 3,
			Want:   "III",
		},
		"4 gets converted to IV": {
			Arabic: 4,
			Want:   "IV",
		},
		"5 gets converted to V": {
			Arabic: 5,
			Want:   "V",
		},
		"6 gets converted to VI": {
			Arabic: 6,
			Want:   "VI",
		},
		"7 gets converted to VII": {
			Arabic: 7,
			Want:   "VII",
		},
		"8 gets converted to VIII": {
			Arabic: 8,
			Want:   "VIII",
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
