package propertybasedtests

import (
	"strings"
)

type RomanNumeral struct {
	value  uint16
	symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1_000, "M"},
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

func convertToRoman(arabic uint16) string {
	var result strings.Builder
	for _, number := range allRomanNumerals {
		for arabic >= number.value {
			result.WriteString(number.symbol)
			arabic -= number.value
		}
	}
	return result.String()
}

func convertToArabic(roman string) uint16 {
	var result uint16 = 0
	for _, number := range allRomanNumerals {
		for strings.HasPrefix(roman, number.symbol) {
			result += number.value
			roman = strings.TrimPrefix(roman, number.symbol)
		}
	}
	return result
}
