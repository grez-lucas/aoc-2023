package day3

import (
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
