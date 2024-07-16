package day2

import (
	"regexp"
	"strconv"
)

type GameSet struct {
  Red  int
  Green  int
  Blue  int
}

type Game struct {
  id int
  gameSets []GameSet
} 

func newGameSet(colors ...int) *GameSet {
  red, green, blue := 0, 0, 0 // default values

  if len(colors) > 0 { red  = colors[0] }
  if len(colors) > 1 { green  = colors[0] }
  if len(colors) > 2 { blue  = colors[0] }

  g := GameSet{Red: red, Green: green, Blue: blue}
  return &g
}

func newGame(id int) *Game {
  g := Game{
    id: id,
    gameSets: []GameSet{},
  }

  return &g
}

func parseGame(line string) *Game{
 var id int = ParseId(line)
  // Split string by ";" delimeter, and parse sub strings with parseGameSet
  

  return newGame(id)
}

func ParseId(line string) int {
  re := regexp.MustCompile(`Game\s(\d+):`) 

  matches := re.FindAllStringSubmatch(line, -1)
  for _, m := range matches {
   if len(matches) > 0 {
      gameIDStr := m[1]

      GameID, err := strconv.Atoi(gameIDStr)

      if err != nil { panic(err)}
    
      return GameID
    } 
  }
  

  return -1
}

func IsPossibleGame(game Game,
  redCubes int,
  greenCubes int,
  blueCubes int) bool {
  
  sumRed := 0
  sumGreen := 0
  sumBlue := 0

  for _, gameSet := range game.gameSets{ 
    sumRed += gameSet.Green
    sumGreen += gameSet.Green
    sumBlue += gameSet.Blue
  } 
  
  return ( redCubes >= sumRed && greenCubes >= sumGreen && blueCubes >= sumBlue)

}
