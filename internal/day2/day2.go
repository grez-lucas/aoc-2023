package day2

import (
	"fmt"
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

var (
    redPattern   = regexp.MustCompile(`(\d+)\sred`)
    greenPattern = regexp.MustCompile(`(\d+)\sgreen`)
    bluePattern  = regexp.MustCompile(`(\d+)\sblue`)
)

func newGameSet(colors ...int) *GameSet {
  red, green, blue := 0, 0, 0 // default values

  if len(colors) > 0 { red  = colors[0] }
  if len(colors) > 1 { green  = colors[1] }
  if len(colors) > 2 { blue  = colors[2] }

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

func parseGameSets(line string) *[]string {

  result := make([]string, 1)
  gameset_re := regexp.MustCompile(`(\d[\w\s\d,]+)[\n;\z]`)

  matches := gameset_re.FindAllStringSubmatch(line, -1)


  for _, match := range matches {

    if len(matches) > 0 {

      result = append(result, match[1])
    }
    
  }
  return &result
}

func ParseGameSet(gamesetString string) (*GameSet, error) {


  parseColor := func(pattern *regexp.Regexp, name string) (int, error) {
    matches := pattern.FindStringSubmatch(gamesetString)

    if len(matches) < 2 {
      return 0, nil
    }
    value, err := strconv.Atoi(matches[1])
    if err != nil {
      return 0, fmt.Errorf("Invalid %s value %v", name, err)
    }
    return value, nil
  }

  red, err := parseColor(redPattern, "red")
  if err != nil {
    return nil, err
  }

  green, err := parseColor(greenPattern, "green")
  if err != nil {
    return nil, err
  }

  blue, err := parseColor(bluePattern, "blue")
  if err != nil {
    return nil, err
  }

  // Parse the line and identify gameSet
  gs := newGameSet(red, green, blue)

  return gs, nil

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
