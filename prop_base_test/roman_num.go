package main

import (
	"strings"
)

func ConvertToRoman(num int) string {
	var result strings.Builder

	for num > 0 {
		switch {
		case num > 4:
			result.WriteString("V")
			num -= 5
		case num > 3:
			result.WriteString("IV")
			num -= 4
		default:
			result.WriteString("I")
			num -= 1
		}
	}

	return result.String()
}
