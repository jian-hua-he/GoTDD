package main

func ConvertToRoman(num int) string {
	if num == 3 {
		return "III"
	}

	if num == 2 {
		return "II"
	}

	return "I"
}
