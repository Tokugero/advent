package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func main() {
	utils.GetInput()
	input, err := utils.ReadInput()

	if err != nil {
		panic(err)
	}

	// final values used to calculate results
	var digitValues []int
	var parsedValues []int

	// static map of digits used in parsing
	wordDigits := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// Parse input line for numbers for solution 1 & 2
	for _, line := range input {
		// Placeholder variables to transform
		stitchingLine := ""
		poppedLine := line

		// Iterate through wordDigits map and replace words with digits for solution 2
		for word, digit := range wordDigits {
			if strings.Contains(poppedLine, word) {
				splitLine := strings.Split(poppedLine, word)
				for i, split := range splitLine {
					if i < len(splitLine)-1 {
						// may tech jesus have mercy on my soul, leave the surrounding letters in case they cascade to other
						// spelled numbers)
						stitchingLine += split + string(word[0]) + digit + string(word[len(word)-1])
					} else if i == len(splitLine)-1 {
						stitchingLine += split
					}
				}
				poppedLine = stitchingLine
			}
			stitchingLine = ""
		}

		// With the original list and the parsed list, find the first and last number and add them together
		digitalNumbers := findNumbers(line)
		parsedNumbers := findNumbers(poppedLine)

		if len(digitalNumbers) >= 1 {
			calibrationValue, _ := strconv.Atoi(digitalNumbers[0] + digitalNumbers[len(digitalNumbers)-1])
			digitValues = append(digitValues, calibrationValue)
		}

		if len(parsedNumbers) >= 1 {
			calibrationValue, _ := strconv.Atoi(parsedNumbers[0] + parsedNumbers[len(parsedNumbers)-1])
			parsedValues = append(parsedValues, calibrationValue)
		}
	}

	fmt.Println("Part 1: ", utils.Sum(digitValues))
	fmt.Println("Part 2: ", utils.Sum(parsedValues))
}

func findNumbers(line string) []string {
	var numbers []string
	for _, char := range line {
		if _, err := strconv.Atoi(string(char)); err == nil {
			numbers = append(numbers, string(char))
		}
	}
	return numbers
}
