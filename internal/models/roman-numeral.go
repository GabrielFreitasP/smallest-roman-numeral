package models

// Roman numeral search model
type RomanNumeralSearch struct {
	Text string `json:"text"`
}

// Roman numeral model
type RomanNumeral struct {
	Number string `json:"number"`
	Value  int    `json:"value"`
}
