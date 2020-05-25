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

	for i := 0; i < len(roman); i += 1 {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]
			value := allRomanNum.ValueOf(symbol, nextSymbol)
			if value != 0 {
				result += value
				i += 1
			} else {
				result += allRomanNum.ValueOf(symbol)
			}
		} else {
			result += allRomanNum.ValueOf(symbol)
		}
	}

	return result
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' ||
		currentSymbol == 'X' ||
		currentSymbol == 'C'

	return index+1 < len(roman) && isSubtractiveSymbol
}
