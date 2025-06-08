package propertybasedtests

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 19, Roman: "XIX"},
	{Arabic: 28, Roman: "XXVIII"},
	{Arabic: 30, Roman: "XXX"},
	{Arabic: 31, Roman: "XXXI"},
	{Arabic: 35, Roman: "XXXV"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 55, Roman: "LV"},
	{Arabic: 80, Roman: "LXXX"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 120, Roman: "CXX"},
	{Arabic: 488, Roman: "CDLXXXVIII"},
	{Arabic: 612, Roman: "DCXII"},
	{Arabic: 932, Roman: "CMXXXII"},
	{Arabic: 1049, Roman: "MXLIX"},
	{Arabic: 1924, Roman: "MCMXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
}

func TestConvertArabicToRoman(t *testing.T) {
	for _, testCase := range cases {
		name := fmt.Sprintf("%d should convert to roman %q", testCase.Arabic, testCase.Roman)
		t.Run(name, func(t *testing.T) {
			got := convertToRoman(testCase.Arabic)
			if got != testCase.Roman {
				t.Errorf("got %q, want %q for converting Arabic: %d", got, testCase.Roman, testCase.Arabic)
			}
		})
	}
}

func TestConvertRomanToArabic(t *testing.T) {
	for _, testCase := range cases {
		name := fmt.Sprintf("%q should convert to arabic %d", testCase.Roman, testCase.Arabic)
		t.Run(name, func(t *testing.T) {
			got := convertToArabic(testCase.Roman)
			if got != testCase.Arabic {
				t.Errorf("got %d, want %d for converting Roman: %q", got, testCase.Arabic, testCase.Roman)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		if arabic < 0 || arabic > 3999 {
			log.Println(arabic)
			return true
		}
		t.Log("testing", arabic)
		roman := convertToRoman(arabic)
		fromRoman := convertToArabic(roman)
		return arabic == fromRoman
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}
