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
  var totalSum int = 0


  for scanner.Scan() {
    line := scanner.Text()
    err := scanner.Err()

    if err != nil {
      panic(err)
    }
    // Check if the game is valid, if so add it to the totalSum
    game, err := day2.ParseGame(line)

    if err != nil {
      panic ( fmt.Errorf("Error parsing game: %v", err))
    }


    if day2.IsPossibleGame(*game, 12, 13, 14) { 

      fmt.Printf("Game %v is possible with 12R, 13G, 14B !\n", game)
      totalSum+= day2.GetId(*game)
    } else {

     // fmt.Printf("Game %v is NOT possible with 12R, 13G, 14B !\n", game)
    }
  }


  fmt.Println("Result: ", totalSum)
}
