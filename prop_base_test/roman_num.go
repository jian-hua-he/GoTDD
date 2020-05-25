package main

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}

	return false
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(num int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	result := 0
	allRomanNum := RomanNumerals(allRomanNumerals)

	for _, symbols := range windowedRoman(roman).Symbols() {
		result += allRomanNum.ValueOf(symbols...)
	}

	return result
}

type windowedRoman string

func (w windowedRoman) Symbols() [][]byte {
	var symbols [][]byte
	allRomanNum := RomanNumerals(allRomanNumerals)

	for i := 0; i < len(w); i += 1 {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNum.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{byte(symbol), byte(w[i+1])})
			i += 1
		} else {
			symbols = append(symbols, []byte{byte(symbol)})
		}
	}

	return symbols
}

func isSubtractive(symbol uint8) bool {
	symbols := []uint8{'I', 'X', 'C'}
	for _, s := range symbols {
		if symbol == s {
			return true
		}
	}

	return false
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' ||
		currentSymbol == 'X' ||
		currentSymbol == 'C'

	return index+1 < len(roman) && isSubtractiveSymbol
}
