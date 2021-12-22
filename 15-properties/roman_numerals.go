package main

import (
	"strings"
)

var romanLiterals = []struct {
	Symbol string
	Value  int
}{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"D", 400},
	{"C", 100},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ConvertToArabic(roman string) (arabic int) {
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

func ConvertToRoman(arabic int) (roman string) {

	for _, literal := range romanLiterals {
		count := arabic / literal.Value
		arabic -= count * literal.Value
		roman += strings.Repeat(literal.Symbol, count)
	}
	return

}
