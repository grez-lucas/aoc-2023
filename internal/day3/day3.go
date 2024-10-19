package day3

import (
	"slices"
)

type position struct {
	r int
	c int
}

func parseTabletLine(row int, line string) (map[position]byte, map[position]byte) {
	// Gets a Tablet string, returns a NumbersMap and a SymbolsMap

	// Map coordinates to the char
	// Char could be '1234567890.'

	numbersMap := make(map[position]byte, 140)
	symbolsMap := make(map[position]byte, 140)

	for col, character := range line {
		byteCharacter := byte(character)

		if isSymbol(byteCharacter) {
			symbolsMap[position{row, col}] = byteCharacter
		}

		if isNumber(byteCharacter) {
			numbersMap[position{row, col}] = byteCharacter
		}

	}

	return numbersMap, symbolsMap

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
