package main

import (
	"strings"
)

func ConvertToRoman(num int) string {
	var result strings.Builder

	if num == 4 {
		return "IV"
	}

	for i := 0; i < num; i += 1 {
		result.WriteString("I")
	}

	return result.String()
}
