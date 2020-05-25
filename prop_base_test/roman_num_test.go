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
			Want:   "XL",
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
		"90 gets converted to XC": {
			Arabic: 90,
			Want:   "XC",
		},
		"100 gets converted to C": {
			Arabic: 100,
			Want:   "C",
		},
		"400 gets converted to CD": {
			Arabic: 400,
			Want:   "CD",
		},
		"500 gets converted to D": {
			Arabic: 500,
			Want:   "D",
		},
		"900 gets converted to CM": {
			Arabic: 900,
			Want:   "CM",
		},
		"1000 gets converted to M": {
			Arabic: 1000,
			Want:   "M",
		},
		"1984 gets converted to MCMLXXXIV": {
			Arabic: 1984,
			Want:   "MCMLXXXIV",
		},
		"3999 gets converted to MMMCMXCIX": {
			Arabic: 3999,
			Want:   "MMMCMXCIX",
		},
		"2014 gets converted to MMXIV": {
			Arabic: 2014,
			Want:   "MMXIV",
		},
		"1006 gets converted to MVI": {
			Arabic: 1006,
			Want:   "MVI",
		},
		"798 gets converted to DCCXCVIII": {
			Arabic: 798,
			Want:   "DCCXCVIII",
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
