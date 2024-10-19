package main

import (
	"bufio"
	"fmt"
	"github.com/grez-lucas/aoc-2023/utils"
	"os"
)

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

		// Combine line's border digits and add them to the total sum
		res, err := utils.CombineBorderDigits(line)
		utils.CheckError(err)

		totalSum += res
	}

	utils.CheckError(err)

	fmt.Println("Result: ", totalSum)
}
