package main

import (
  "bufio"
  "os"
  "github.com/grez-lucas/aoc-2023/utils"
  "fmt"
)

type GameSet struct {
  Red  int
  Green  int
  Blue  int
}

func newGameSet(colors ...int) *GameSet {
  red, green, blue := 0, 0, 0 // default values

  if len(colors) > 0 { red  = colors[0] }
  if len(colors) > 1 { green  = colors[0] }
  if len(colors) > 2 { blue  = colors[0] }

  g := GameSet{Red: red, Green: green, Blue: blue}
  return &g
}

func main() {

  file, err := os.Open("./input.txt")
  
  utils.CheckError(err)
  defer file.Close() // Make sure the file closes regardless of how main() ends

  scanner := bufio.NewScanner(file)
  var totalSum int = 0


  for scanner.Scan() {
    line := scanner.Text()
    err := scanner.Err()
    utils.CheckError(err)

    // Check if the game is valid, if so add it to the totalSum
    res, err := utils.CombineBorderDigits(line)
    utils.CheckError(err)

    totalSum+=res
  }

  utils.CheckError(err)

  fmt.Println("Result: ", totalSum)
}
