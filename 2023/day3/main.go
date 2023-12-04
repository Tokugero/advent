package main

import (
	"fmt"
	"regexp"
	"strconv"
	utils "utils"
)

type Part struct {
	number   int
	firstPos int
	lastPos  int
}

type Cog struct {
	numbers  []int
	multiple int
	firstPos int
	lastPos  int
}

func main() {
	utils.GetInput()
	input, err := utils.ReadInput()
	if err != nil {
		panic(err)
	}

	parts := []Part{}
	cogs := []Cog{}

	digitRegex := regexp.MustCompile(`\d+`)
	specialRegex := regexp.MustCompile(`[\!\@\#\$\%\^\&\*\(\)\_\+\-\=\[\]\{\}\;\:\'\"\,\/\<\>\?]`)
	cogsRegex := regexp.MustCompile(`\*`)

	for i, line := range input {
		if line == "" {
			continue
		}

		// parse the string for numbers, record positions
		potentialParts := digitRegex.FindAllStringIndex(line, -1)
		potentialCogs := cogsRegex.FindAllStringIndex(line, -1)

		// check around indexes for symbols, above, below, left and right +/-1 for symbol not period
		for _, potentialPart := range potentialParts {
			fPos := potentialPart[0]
			lPos := potentialPart[1]

			ppNumber, _ := strconv.Atoi(line[fPos:lPos])

			// check above
			if i > 0 {
				if lookAround(&input, i, specialRegex, fPos, lPos, "up", 1) {
					parts = append(parts, Part{number: ppNumber, firstPos: fPos, lastPos: lPos})
				}
			}
			// check below
			if i <= len(input) {
				if lookAround(&input, i, specialRegex, fPos, lPos, "down", 1) {
					parts = append(parts, Part{number: ppNumber, firstPos: fPos, lastPos: lPos})
				}
			}
			// check left or right
			if lookAround(&input, i, specialRegex, fPos, lPos, "side", 1) {
				parts = append(parts, Part{number: ppNumber, firstPos: fPos, lastPos: lPos})
			}
		}

		for _, potentialCog := range potentialCogs {
			fPos := potentialCog[0]
			lPos := potentialCog[1]

			cog := Cog{numbers: []int{}, multiple: 0, firstPos: fPos, lastPos: lPos}

			// check above
			if i > 0 {
				if lookAround(&input, i, digitRegex, fPos, lPos, "up", 1) {
					potentialCogNumbers := digitRegex.FindAllStringIndex(input[i-1], -1)
					for _, potentialCogNumber := range potentialCogNumbers {
						if (fPos >= potentialCogNumber[0] && fPos <= potentialCogNumber[1]) || (lPos >= potentialCogNumber[0] && lPos <= potentialCogNumber[1]) {
							cogNumber, _ := strconv.Atoi(input[i-1][potentialCogNumber[0]:potentialCogNumber[1]])
							cog.numbers = append(cog.numbers, cogNumber)
						}
					}
				}
			}
			// check below
			if i <= len(input) {
				if lookAround(&input, i, digitRegex, fPos, lPos, "down", 1) {
					potentialCogNumbers := digitRegex.FindAllStringIndex(input[i+1], -1)
					for _, potentialCogNumber := range potentialCogNumbers {
						if (fPos >= potentialCogNumber[0] && fPos <= potentialCogNumber[1]) || (lPos >= potentialCogNumber[0] && lPos <= potentialCogNumber[1]) {
							cogNumber, _ := strconv.Atoi(input[i+1][potentialCogNumber[0]:potentialCogNumber[1]])
							cog.numbers = append(cog.numbers, cogNumber)
						}
					}
				}
			}
			// check left or right
			if lookAround(&input, i, digitRegex, fPos, lPos, "side", 1) {
				potentialCogNumbers := digitRegex.FindAllStringIndex(line, -1)
				for _, potentialCogNumber := range potentialCogNumbers {
					if (fPos >= potentialCogNumber[0] && fPos <= potentialCogNumber[1]) || (lPos >= potentialCogNumber[0] && lPos <= potentialCogNumber[1]) {
						cogNumber, _ := strconv.Atoi(input[i][potentialCogNumber[0]:potentialCogNumber[1]])
						cog.numbers = append(cog.numbers, cogNumber)
					}
				}
			}

			if len(cog.numbers) >= 2 {
				cogs = append(cogs, cog)
			}
		}
	}

	sumPart1 := 0
	for _, part := range parts {
		sumPart1 += part.number
	}

	sumPart2 := 0
	for _, cog := range cogs {
		multiple := 1
		for _, number := range cog.numbers {
			multiple *= number
		}
		sumPart2 += multiple
	}

	fmt.Println("Part 1: ", sumPart1)
	fmt.Println("Part 2: ", sumPart2)
}

func lookAround(input *[]string, line int, re *regexp.Regexp, fPos int, lPos int, direction string, depth int) bool {
	modX := 0

	switch direction {
	case "up":
		modX = -depth
	case "down":
		modX = depth
	}

	if (*input)[line+modX] == "" {
		return false
	}

	if fPos >= depth && lPos < len((*input)[line+modX]) {
		if re.MatchString((*input)[line+modX][fPos-depth : lPos+depth]) {
			return true
		}
	} else if fPos > depth {
		if re.MatchString((*input)[line+modX][fPos-depth : lPos]) {
			return true
		}
	} else if lPos < len((*input)[line+modX]) {
		if re.MatchString((*input)[line+modX][fPos : lPos+depth]) {
			return true
		}
	}

	return false
}
