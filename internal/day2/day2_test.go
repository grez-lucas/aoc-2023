package day2_test

import (
	"slices"
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
      {
        name:  "Game 3 has Id 3",
        input: "Game 3: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
        expected: 3,
        },
      {
        name:  "Game 22 has Id 22",
        input: "Game 22: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
        expected: 22,
        },
      }

  	for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result := day2.ParseId(tt.input)

      // Log matches for debugging
      t.Logf("Found id: %d", result)

      // Assert
      if result != tt.expected {
        t.Errorf("Test %s: Got result %d, expected %d", tt.name, result, tt.expected)
        return
	}


})
  } 
}

func TestParseGameSet(t *testing.T) {

  // Arrange
  var tests = []struct {
    name string
    input string
    expected day2.GameSet
  } {
{
        name:  "parses standard gameSet",
        input: "3 green, 4 blue, 1 red;",
        expected: day2.GameSet{ Red: 1, Green: 3, Blue: 4},
        },
      {
        name:  "parses GameSet with missing values",
        input: "1 blue, 2 green",
        expected: day2.GameSet{ Red: 0, Green: 2, Blue: 1},
        },
      }

  	for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result, err := day2.ParseGameSet(tt.input)

      // Log matches for debugging
      // t.Logf("Found id: %d", result)

      // Assert
      if err != nil {
        t.Errorf("Test %s: Raised an error %v", tt.name, err)
      }

      if *result != tt.expected {
        t.Errorf("Test %s: Got result %d, expected %d", tt.name, result, tt.expected)
        return
	}


})
  } 
}

func TestParseGameSets(t *testing.T) {

  // Arrange
  var tests = []struct {
    name string
    input string
    expected []string
  } {
      {
        name:  "parses standard GameSets",
        input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
        expected: []string{"3 blue, 4 red;", "1 red, 2 green, 6 blue;", "2 green"},
      },
      {
        name:  "parses multi-line GameSets",
        input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
        expected: []string{"3 blue, 4 red;", "1 red, 2 green, 6 blue;", "2 green", "1 blue, 2 green;", "3 green, 4 blue, 1 red;", "1 green, 1 blue"},
      },
    }

  	for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result := day2.ParseGameSets(tt.input)

      // Log matches for debugging
      // t.Logf("Found id: %d", result)

      // Assert

      if slices.Equal(*result, tt.expected){
        t.Errorf("Test %s: Got result %v, expected %v", tt.name, result, tt.expected)
        return
	}


})
  } 
}
