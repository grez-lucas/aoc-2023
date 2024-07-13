package utils_test

import (
  "testing"
  "github.com/grez-lucas/aoc-2023/utils"
)

func TestCombineBorderDigits_EmptyInput(t *testing.T) {
  // Test case: Empty string input
  // Arrange
  input := ""
  // Act
  result, err := utils.CombineBorderDigits(input)
  // Assert
  expected := 0

  if err != nil {
    t.Errorf("Error: %s", err.Error())
  }

  if result != expected {
    t.Errorf("Got %d, expected %d", result, expected)
  }
}

func TestFindDigitStrings_TableDriven(t *testing.T) {
  // Arrange
  var tests = []struct {
    name string
    input string
    expected []utils.DigitFound
  } {

      {
        name:  "eightwothree should have 3 matches",
        input: "eightwothree",
        expected: []utils.DigitFound{
          {Value: 8, Index: 0},
          {Value: 2, Index: 4},
          {Value: 3, Index: 7},
        },
      },
      {
        name:  "eightwothree should have 3 matches",
        input: "eightwothree",
        expected: []utils.DigitFound{
          {Value: 8, Index: 0},
          {Value: 2, Index: 4},
          {Value: 3, Index: 7},
        },
      }, 

    }

  	for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result := utils.FindDigitStrings(tt.input)

      // Log matches for debugging
      t.Logf("Matches found: %v", result)

      // Assert
      if len(result) != len(tt.expected) {
        t.Errorf("Test %s: Got %d results, expected %d", tt.name, len(result), len(tt.expected))
        return
      }

      for i, expected := range tt.expected {
        if result[i].Value != expected.Value || result[i].Index != expected.Index {
          t.Errorf("Test %s: Mismatch at index %d. Got %v, expected %v", tt.name, i, result[i], expected)
        }
      }
    })
	}



}
