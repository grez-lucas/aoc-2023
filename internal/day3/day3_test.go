package day3

import (
	"maps"
	"testing"
)

func Test_isSymbol(t *testing.T) {
	// Arrange
	var tests = []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "Symbol is symbol",
			input:    byte('*'),
			expected: true,
		},
		{
			name:     "Non-Symbol is not symbol",
			input:    byte('8'),
			expected: false,
		},
		{
			name:     "Non-Symbol is not symbol",
			input:    byte('A'),
			expected: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			// Act
			result := isSymbol(tt.input)

			// Assert
			if result != tt.expected {
				t.Errorf("Incorrect result, expected `%v` got `%v`", tt.expected, result)
			}

		})

	}
}
func Test_isNumber(t *testing.T) {
	// Arrange
	var tests = []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "Number is number",
			input:    byte('9'),
			expected: true,
		},
		{
			name:     "Non-Number is not number",
			input:    byte('*'),
			expected: false,
		},
		{
			name:     "Whitespace is not number",
			input:    byte(' '),
			expected: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			// Act
			result := isNumber(tt.input)

			// Assert
			if result != tt.expected {
				t.Errorf("Incorrect result, expected `%v` got `%v`", tt.expected, result)
			}

		})

	}
}

func TestParseTabletLine(t *testing.T) {
	// Arrange
	var tests = []struct {
		name               string
		inputRow           int
		inputString        string
		expectedNumbersMap map[Position]byte
		expectedSymbolsMap map[Position]byte
	}{
		{
			name:        "Normal inputString",
			inputRow:    8,
			inputString: "...45....888",
			expectedNumbersMap: map[Position]byte{
				{8, 3}:  byte('4'),
				{8, 4}:  byte('5'),
				{8, 9}:  byte('8'),
				{8, 10}: byte('8'),
				{8, 11}: byte('8'),
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			// Act
			resNumbersMap, resSymbolsMap := parseTabletLine(
				tt.inputRow,
				tt.inputString,
			)

			// Assert
			if !maps.Equal(resNumbersMap, tt.expectedNumbersMap) {
				t.Errorf("Incorrect result, expected `%v` got `%v`", tt.expectedNumbersMap, resNumbersMap)
			}

			if !maps.Equal(resSymbolsMap, tt.expectedSymbolsMap) {
				t.Errorf("Incorrect result, expected `%v` got `%v`", tt.expectedSymbolsMap, resSymbolsMap)
			}

		})

	}
}

func Test_isPartNumber(t *testing.T) {
	// Arrange
	var tests = []struct {
		name            string
		inputPosition   Position
		inputSymbolsMap map[Position]byte
		expected        bool
	}{
		{
			name:          "Symbol on top",
			inputPosition: Position{9, 3},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: true},
		{
			name:          "Symbol on bottom",
			inputPosition: Position{7, 3},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: true},
		{
			name:          "Symbol on left",
			inputPosition: Position{8, 4},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: true},
		{
			name:          "Symbol on right",
			inputPosition: Position{8, 2},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: true},
		{
			name:          "Symbol on top-left",
			inputPosition: Position{9, 4},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: true},
		{
			name:          "Symbol on top-right",
			inputPosition: Position{9, 2},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: true},
		{
			name:          "Symbol on bottom-right",
			inputPosition: Position{7, 2},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: true},
		{
			name:          "Symbol on bottom-left",
			inputPosition: Position{7, 4},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: true},
		{
			name:          "Symbol on same Position",
			inputPosition: Position{8, 3},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: false},
		{
			name:          "Symbol on non-adjacent Position",
			inputPosition: Position{40, 40},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: false},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			// Act

			result := isPartNumber(tt.inputPosition, tt.inputSymbolsMap)

			if result != tt.expected {
				t.Errorf("Incorrect result, expected `%v` got `%v`", tt.expected, result)
			}
		})
	}
}

func Test_sumTotalPartNumbers(t *testing.T) {
	// Arrange
	var tests = []struct {
		name            string
		inputNumbersMap map[Position]byte
		inputSymbolsMap map[Position]byte
		expected        int
	}{

		{name: "Normal case",
			inputNumbersMap: map[Position]byte{
				{8, 2}: byte('9'),
			},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: 9,
		},
		{name: "Multiple adjacent numbers",
			inputNumbersMap: map[Position]byte{
				{8, 2}: byte('9'),
				{8, 4}: byte('9'),
				{7, 2}: byte('9'),
				{7, 4}: byte('9'),
			},
			inputSymbolsMap: map[Position]byte{
				{8, 3}: byte('*'),
			},
			expected: 36,
		},
		{name: "No adjacent numbers",
			inputNumbersMap: map[Position]byte{
				{8, 2}: byte('9'),
				{8, 4}: byte('9'),
				{7, 2}: byte('9'),
				{7, 4}: byte('9'),
			},
			inputSymbolsMap: map[Position]byte{
				{2, 3}: byte('*'),
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Act
			result, err := sumTotalPartNumbers(tt.inputNumbersMap, tt.inputSymbolsMap)

			// Assert
			if err != nil {
				t.Errorf("Got error, `%v`", err)
			}

			if result != tt.expected {
				t.Errorf("Inocrrect result, expected `%v`, got `%v`", tt.expected, result)
			}
		})
	}

}
