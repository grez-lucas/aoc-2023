package main

import (
  "bufio"
  "os"
  "github.com/grez-lucas/aoc-2023/internal/day2"
  "fmt"
)


func main() {

  file, err := os.Open("./input.txt")
  
  if err != nil {
    panic (fmt.Errorf("Error opening file: %v", err))
  }
  defer file.Close() // Make sure the file closes regardless of how main() ends

  scanner := bufio.NewScanner(file)
  var validGameSum int = 0
  var gameSetPowerSum int = 0


  for scanner.Scan() {
    line := scanner.Text()
    err := scanner.Err()

    if err != nil {
      panic(err)
    }
    // Check if the game is valid, if so add it to the validGameSum
    game, err := day2.ParseGame(line)

    if err != nil {
      panic ( fmt.Errorf("Error parsing game: %v", err))
    }


    if day2.IsPossibleGame(*game, 12, 13, 14) { 

      fmt.Printf("Game %v is possible with 12R, 13G, 14B !\n", game)
      validGameSum+= day2.GetId(*game)
    } else {

     // fmt.Printf("Game %v is NOT possible with 12R, 13G, 14B !\n", game)
    }

    // Get the power of the min cube GameSet for that game and add it up
    gameSetPowerSum += day2.GetGameSetPower(day2.GetMinGameCubes(*game))
  }


  fmt.Println("ValidGame Sum Result: ", validGameSum)
  fmt.Println("GameSetPowerSum Result: ", gameSetPowerSum)
}
