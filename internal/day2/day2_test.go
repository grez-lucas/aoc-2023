package day2

import (
	"slices"
	"testing"

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
      result := ParseId(tt.input)

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
    expected GameSet
  } {
{
        name:  "parses standard gameSet",
        input: "3 green, 4 blue, 1 red;",
        expected: GameSet{ Red: 1, Green: 3, Blue: 4},
        },
      {
        name:  "parses GameSet with missing values",
        input: "1 blue, 2 green",
        expected: GameSet{ Red: 0, Green: 2, Blue: 1},
        },
      }

  	for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result, err := ParseGameSet(tt.input)

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
        name:  "parses single GameSet",
        input: "Game 1: 3 blue, 4 red",
        expected: []string{"3 blue, 4 red"},
      },
      {
        name:  "parses multiple GameSets",
        input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
        expected: []string{"3 blue, 4 red;", "1 red, 2 green, 6 blue;", "2 green"},
      },
    }

  	for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result := ParseGameSets(tt.input)

      // Log matches for debugging
      // t.Logf("Found id: %d", result)

      // Assert

      if !slices.Equal(result, tt.expected) {
        t.Errorf("Test %s: Got result %v, expected %v", tt.name, result, tt.expected)
        return
	}


})
  } 
}

func TestParseGame(t *testing.T) {

  // Arrange
  var tests = []struct {
    name string
    input string
    expected Game
  } {
      {
        name:  "parses multiple GameSet standard Game",
        input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
        expected: Game{ 
          id: 1,
          gameSets: []GameSet{ 
            {Red: 4, Green: 0, Blue: 3}, 
            {Red: 1, Green: 2, Blue: 6}, 
            {Red: 0, Green: 2, Blue: 0}, 
          }, 
        },
      },
      {
        name:  "parses single GameSet GameSets",
        input: "Game 2: 1 blue, 2 green",
        expected: Game{ id: 2, gameSets: []GameSet{
          {Red: 0, Green: 2, Blue: 1},
        },
        },
      },
    }

  	for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result, err := ParseGame(tt.input)

      // Log matches for debugging
      // t.Logf("Found id: %d", result)
      if err != nil {

        t.Errorf("Test %s: Got unexpected error %v", tt.name, err)
      }

      // Assert
      if result.id != tt.expected.id {
        t.Errorf("Test %s: Got id %d, expected %d", tt.name, result.id, tt.expected.id)
      }

      if !slices.Equal(result.gameSets, tt.expected.gameSets) {
        t.Errorf("Test %s: Got result %v, expected %v", tt.name, result, tt.expected)
        return
	}


})
  } 
}

func TestIsPossibleGame(t *testing.T) {

  // Arrange
  var tests = []struct {
    name string
    input Game
    expected bool
  } {
      {
        name:  "possible game is possible",
        input: Game{
          id: 1,
          gameSets: []GameSet{ 
            {Red: 4, Green: 0, Blue: 3}, 
            {Red: 1, Green: 2, Blue: 6}, 
            {Red: 0, Green: 2, Blue: 0}, 
          }, 
        },
        expected: true,
      },
      {
        name:  "impossible game is not possible",
        input: Game{
          id: 1,
          gameSets: []GameSet{ 
            {Red: 4, Green: 1600, Blue: 3}, 
            {Red: 1, Green: 2, Blue: 6}, 
            {Red: 1, Green: 2, Blue: 0}, 
          }, 
        },
        expected: false,
      },
      {
        name:  "possible game with less than enough cubes is possible",
        input: Game{
          id: 1,
          gameSets: []GameSet{ 
            {Red: 0, Green: 0, Blue: 0}, 
            {Red: 1, Green: 0, Blue: 6}, 
            {Red: 0, Green: 2, Blue: 0}, 
          }, 
        },
        expected: true,
      },
    }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result := IsPossibleGame(tt.input, 5, 4, 9)

      // Assert

      if result != tt.expected{
        t.Errorf("Test %s: Got result %v, expected %v", tt.name, result, tt.expected)
        return
      }

    })
  } 
}
func TestGetMinGameCubes(t *testing.T) {

  // Arrange
  var tests = []struct {
    name string
    input Game
    expected GameSet
  } {
      {
        name:  "Returns correct gameset",
        input: Game{
          id: 1,
          gameSets: []GameSet{ 
            {Red: 4, Green: 0, Blue: 3}, 
            {Red: 1, Green: 2, Blue: 6}, 
            {Red: 0, Green: 2, Blue: 0}, 
          }, 
        },
        expected: GameSet{Red: 4, Green: 2, Blue: 6},
      },
      {
        name:  "Returns correct gameset",
        input: Game{
          id: 2,
          gameSets: []GameSet{ 
            {Red: 0, Green: 2, Blue: 1}, 
            {Red: 1, Green: 3, Blue: 4}, 
            {Red: 0, Green: 1, Blue: 1}, 
          }, 
        },
        expected: GameSet{Red: 1, Green: 3, Blue: 4},
      },
      {
        name:  "Returns correct gameset",
        input: Game{
          id: 3,
          gameSets: []GameSet{ 
            {Red: 20, Green: 8, Blue: 6}, 
            {Red: 4, Green: 13, Blue: 5}, 
            {Red: 1, Green: 5, Blue: 0}, 
          }, 
        },
        expected: GameSet{Red: 20, Green: 13, Blue: 6},
      },
      {
        name:  "Returns correct gameset",
        input: Game{
          id: 4,
          gameSets: []GameSet{ 
            {Red: 3, Green: 1, Blue: 6}, 
            {Red: 6, Green: 3, Blue: 0}, 
            {Red: 14, Green: 3, Blue: 15}, 
          }, 
        },
        expected: GameSet{Red: 14, Green: 3, Blue: 15},
      },
      {
        name:  "Returns correct gameset",
        input: Game{
          id: 5,
          gameSets: []GameSet{ 
            {Red: 6, Green: 3, Blue: 1}, 
            {Red: 1, Green: 2, Blue: 2}, 
          }, 
        },
        expected: GameSet{Red: 6, Green: 3, Blue: 2},
      },
    }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result := GetMinGameCubes(tt.input)

      // Assert

      if result != tt.expected{
        t.Errorf("Test %s: Got result %v, expected %v", tt.name, result, tt.expected)
        return
      }

    })
  } 
}

func TestGetGameSetPower(t *testing.T) {

  // Arrange
  var tests = []struct {
    name string
    input GameSet
    expected int
  } {
      {
        name:  "Returns correct gameset",
        input: GameSet{Red: 4, Green: 0, Blue: 3}, 
        expected: 0,
      },
      {
        name:  "Returns correct gameset",
        input: GameSet{Red: 6, Green: 1, Blue: 1}, 
        expected: 6,
      },
    }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Act
      result := GetGameSetPower(tt.input)

      // Assert

      if result != tt.expected{
        t.Errorf("Test %s: Got result %v, expected %v", tt.name, result, tt.expected)
        return
      }

    })
  } 
}
