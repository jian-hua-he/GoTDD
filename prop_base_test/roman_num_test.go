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
		"9 gets converted to IX": {
			Arabic: 9,
			Want:   "IX",
		},
		"10 gets converted to X": {
			Arabic: 10,
			Want:   "X",
		},
		"14 gets converted to XIV": {
			Arabic: 14,
			Want:   "XIV",
		},
		"18 gets converted to XVIII": {
			Arabic: 18,
			Want:   "XVIII",
		},
		"20 gets converted to XX": {
			Arabic: 20,
			Want:   "XX",
		},
		"39 gets converted to XXXIX": {
			Arabic: 39,
			Want:   "XXXIX",
		},
		"40 gets converted to XL": {
			Arabic: 40,
			Want:   "XXXIX",
		},
		"47 gets converted to XLVII": {
			Arabic: 47,
			Want:   "XLVII",
		},
		"49 gets converted to XLIX": {
			Arabic: 49,
			Want:   "XLIX",
		},
		"50 gets converted to L": {
			Arabic: 50,
			Want:   "L",
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
