package day2_test

import (
  "testing"
  "github.com/grez-lucas/aoc-2023/internal/day2"
)


func TestParseId(t *testing.T) {

  // Arrange
  var tests = []struct {
    name string
    input string
    expected int
  } {
{
        name:  "Game 2 has Id 2",
        input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
        expected: 2,
        },
      }

  	for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result := day2.ParseId(tt.input)

      // Log matches for debugging
      t.Logf("Matches found: %v", result)

      // Assert
      if result != tt.expected {
        t.Errorf("Test %s: Got result %d, expected %d", tt.name, result, tt.expected)
        return
	}



}
  } 
}
