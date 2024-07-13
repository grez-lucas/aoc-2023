package utils

import (
  "fmt" 
  "regexp"
  "strconv"
  "sort"
)

type DigitFound struct {
  Value  int
  Index  int
}

func CheckError(err error) {
  if err != nil {
    panic(err)
  }
}

func FindDigitStrings(word string) ([]DigitFound) {
  
  // Option 2 BETTER: Create a Queue with every number of the string
  // To parse the string: We:
    // Get ALL pattern matches and their index (where they start)
    // Include numbers ("9", "5", etc) in the pattern match
    // Add them all to a queue, in order of index
    // Append them to a string in that exact order.

  var digitStrings = map[string]int{
    "0": 0,
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
  }

  var digitsIndexFound []DigitFound 

  content := []byte(word)
  for key, value := range digitStrings {

    re := regexp.MustCompile(key) 

    matches := re.FindAllIndex(content, -1)

    for _, match := range matches {
      digitFound := DigitFound{
        Value: value,   
        Index: match[0], 
      }

      digitsIndexFound = append(digitsIndexFound, digitFound )

    }
  
  }

  sort.Slice(digitsIndexFound, func(i, j int) bool {
    return digitsIndexFound[i].Index < digitsIndexFound[j].Index
  })
  return digitsIndexFound

}

func joinBorderDigits(digitsIndexFound []DigitFound) (int, error) {
  var str string = fmt.Sprintf("%d%d", digitsIndexFound[0].Value,
    digitsIndexFound[len(digitsIndexFound)-1].Value)
  
  res, err := strconv.Atoi(str) 

  if err != nil {
    return 0, err
  }

  return res, err
}

func CombineBorderDigits(word string) (int, error) {
  fmt.Println("Checking word ", word)

  wordLength := len(word)
  if wordLength < 1 { return 0, nil}

  digitsIndexFound := FindDigitStrings(word)

  res , err := joinBorderDigits(digitsIndexFound)

  return res, err
}


//func CombineBorderDigits(word string) (int, error){
  //fmt.Println("Checking word: ", word)
//
  //wordLength := len(word)
  //if wordLength < 1 { return 0, nil }
//
  //frontIndex := 0
  //backIndex := wordLength - 1
  //var frontDigit rune = '\u0000'
  //var backDigit rune = '\u0000'
//
  //for  frontDigit == '\u0000' || backDigit == '\u0000'{
    //frontChar := rune(word[frontIndex])
    //backChar := rune(word[backIndex])
    //
//
    //if unicode.IsDigit(frontChar) { 
      //frontDigit = frontChar
    //} else { frontIndex++ }
//
    //if unicode.IsDigit(backChar) {
      //backDigit = backChar
    //} else { backIndex--} 
//
    //
  //}
  //res, err := strconv.Atoi(string(frontDigit)+ string(backDigit))
//
  //if err != nil {
    //fmt.Println("Error transforming result string to integer")
  //}
//
  //fmt.Println("Result: ", res)
  //return res, err
  //
//}
