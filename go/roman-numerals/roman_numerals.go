package romannumerals

import "errors"

// Convert Struct
type atr struct {
	ArabicNumber int
	RomanNumber  string
}

var translate = []atr{
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

// ToRomanNumeral converts arabic numbers to Roman
func ToRomanNumeral(i int) (string, error) {
	var output string
	if i < 1 || i > 3000 {
		return "", errors.New("number out of range 1-3000")
	}
	for _, v := range translate {
		for ; i >= v.ArabicNumber; i -= v.ArabicNumber {
			output = output + v.RomanNumber
		}
	}
	return output, nil
}
