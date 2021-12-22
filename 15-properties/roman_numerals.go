package main

import (
	"strings"
)

var romanLiterals = []struct {
	Symbol string
	Value  uint16
}{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ConvertToArabic(roman string) (arabic uint16) {
	for _, literal := range romanLiterals {
		if strings.HasPrefix(roman, literal.Symbol) {
			arabic += literal.Value
			arabic += ConvertToArabic(roman[len(literal.Symbol):])
			roman = ""
			return
		}
	}
	return
}

func ConvertToRoman(arabic uint16) (roman string) {

	for _, literal := range romanLiterals {
		count := arabic / literal.Value
		arabic -= count * literal.Value
		roman += strings.Repeat(literal.Symbol, int(count))
	}
	return

}
