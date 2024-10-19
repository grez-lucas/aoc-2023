package day3

import (
	"slices"
	"strconv"
)

type position struct {
	r int
	c int
}

func sumTotalPartNumbers(
	numbersMap map[position]byte,
	symbolsMap map[position]byte,
) (int, error) {

	sum := 0

	for pos, val := range numbersMap {
		if isPartNumber(pos, symbolsMap) {
			str := string(val)
			intVal, err := strconv.Atoi(str)

			if err != nil {
				return sum, err
			}

			sum = sum + intVal

		}
	}

	return sum, nil

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

func isPartNumber(
	numberPosition position,
	symbolsMap map[position]byte,
) bool {

	// verifies if the number is adjacent to a symbol, even diagonally

	// Iterate through all keys of numbersMap, for each one check if theres an
	// adjacent symbol

	// Check if there's an adjacent key to numberPosition

	for _, diff := range [2]int{-1, 1} {

		// Top & Bottom
		_, ok := symbolsMap[position{numberPosition.r + diff, numberPosition.c}]
		if ok {
			return true
		}

		// Left & Right
		_, ok = symbolsMap[position{numberPosition.r, numberPosition.c + diff}]
		if ok {
			return true
		}

		// Diagonals 1
		_, ok = symbolsMap[position{numberPosition.r + diff, numberPosition.c + diff}]
		if ok {
			return true
		}

		// Diagonals 2
		_, ok = symbolsMap[position{numberPosition.r - diff, numberPosition.c + diff}]
		if ok {
			return true
		}
	}

	return false
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
