package main

import (
	"fmt"
	"strings"
)

func ConvertToRoman(number int) (roman string) {
	romanLiterals := []struct {
		symbol string
		value  int
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

	for _, l := range romanLiterals {
		count := number / l.value
		number -= count * l.value
		roman += strings.Repeat(l.symbol, count)
		//fmt.Println(l)
	}
	return

}

func main() {

	fmt.Println(ConvertToRoman(1984))

}
