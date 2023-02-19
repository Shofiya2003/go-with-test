package roman

import (
	"strings"
)

type RomanNumeral struct {
	value uint16
	roman string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	roman := string(symbols)
	for _, numeral := range r {
		if numeral.roman == roman {
			return numeral.value
		}
	}

	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	roman := string(symbols)
	for _, numeral := range r {
		if numeral.roman == roman {
			return true
		}
	}

	return false
}

var allRomanNumerals = RomanNumerals{
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

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, romanNumeral := range allRomanNumerals {
		for arabic >= romanNumeral.value {
			result.WriteString(romanNumeral.roman)
			arabic -= romanNumeral.value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) (total uint16) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}

	return
}

func isSubtractiveSymbol(currentSymbol uint8) bool {
	return currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractiveSymbol(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}

	}
	return
}
