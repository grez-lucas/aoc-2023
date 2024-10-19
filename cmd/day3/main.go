package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"

	"github.com/grez-lucas/aoc-2023/internal/day3"
)

func main() {
	file, err := os.Open("input.txt")

	defer file.Close()
	if err != nil {
		panic(fmt.Errorf("Error opening input file"))
	}

	numbersMap := make(map[day3.Position]byte, 140)
	symbolsMap := make(map[day3.Position]byte, 140)
	fPos := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parseNumbersMap, parseSymbolsMap := day3.ParseTabletLine(fPos, line)
		fmt.Println(line)

		maps.Copy(parseNumbersMap, numbersMap)
		maps.Copy(parseSymbolsMap, symbolsMap)

		fPos++
	}

	fmt.Println("numbersMap: ", numbersMap)
	fmt.Println("symbolsMap", symbolsMap)

}
