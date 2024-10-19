package day3

import "slices"

type position struct {
	r int
	c int
}

func ParseTablet() {
	// Gets a Tablet string, returns a Map

	// Map coordinates to the char
	// Coordinates could be their own struct
	// Char could be '1234567890.'

	// m := make(map[position]string)

}

func isPartNumber() {
	// verifies if the number is adjacent to a symbol, even diagonally

	// symbols := []string{"+", "%", "$"}
}

func isSymbol(character byte) bool {
	var symbolsString string = "*#$+"
	var symbols []byte = []byte(symbolsString)

	return slices.Contains(symbols, character)

}

func isNumber(character byte) bool {
	var numbersString string = "1234567890"
	var numbers []byte = []byte(numbersString)

	return slices.Contains(numbers, character)
}
