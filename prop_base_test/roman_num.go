package main

import (
	"strings"
)

func ConvertToRoman(num int) string {
	var result strings.Builder

	for i := 0; i < num; i += 1 {
		result.WriteString("I")
	}

	return result.String()
}
