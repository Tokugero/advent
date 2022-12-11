package main

import (
	"fmt"
	"strings"
)

func main() {
	data, _ := ReadInput()

	inputs := [][]string{}

	for _, line := range strings.Split(data, "\n") {
		input := strings.Split(line, " ")
		inputs = append(inputs, input)
	}

	fmt.Println("Part 1: ", calculate(inputs))
	fmt.Println("Part 2: ", calculate(decode(inputs)))
}

func decode(input [][]string) [][]string {
	var decoded [][]string

	for _, round := range input {
		var throw string
		opponent := strings.ToLower(round[0])
		coded := strings.ToLower(round[1])

		if coded == "x" { //lose
			if opponent == "a" { // rock beats sissors
				throw = "z"
			} else if opponent == "b" { // paper beats rock
				throw = "x"
			} else if opponent == "c" { // scissors beats paper
				throw = "y"
			}

		} else if coded == "y" { //draw
			if opponent == "a" {
				throw = "x"
			} else if opponent == "b" {
				throw = "y"
			} else if opponent == "c" {
				throw = "z"
			}

		} else if coded == "z" { //win
			if opponent == "a" { // paper beats rock
				throw = "y"
			} else if opponent == "b" { // scissors beats paper
				throw = "z"
			} else if opponent == "c" { // rock beats sissors
				throw = "x"
			}
		}

		decoded = append(decoded, []string{opponent, throw})
	}

	return decoded
}

func calculate(input [][]string) []int {
	scoreBoard := []int{0, 0}
	score := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"x": 1,
		"y": 2,
		"z": 3,
	}

	for _, round := range input {
		me := strings.ToLower(round[1])
		opponent := strings.ToLower(round[0])

		if won(round) {
			scoreBoard[0] += 6 + score[me]
			scoreBoard[1] += score[opponent]
		} else if lost(round) {
			scoreBoard[0] += score[me]
			scoreBoard[1] += 6 + score[opponent]
		} else if drew(round) {
			scoreBoard[0] += 3 + score[me]
			scoreBoard[1] += 3 + score[opponent]
		} else {
			fmt.Println("Error")
		}
	}

	return scoreBoard
}

func won(input []string) bool {
	shifted := convert(input)

	if shifted[0] == "a" && shifted[1] == "b" {
		return true
	} else if shifted[0] == "b" && shifted[1] == "c" {
		return true
	} else if shifted[0] == "c" && shifted[1] == "a" {
		return true
	} else {
		return false
	}
}

func lost(input []string) bool {
	shifted := convert(input)

	if shifted[0] == "a" && shifted[1] == "c" {
		return true
	} else if shifted[0] == "b" && shifted[1] == "a" {
		return true
	} else if shifted[0] == "c" && shifted[1] == "b" {
		return true
	} else {
		return false
	}
}

func drew(input []string) bool {
	shifted := convert(input)

	if shifted[0] == shifted[1] {
		return true
	} else {
		return false
	}
}

func convert(input []string) []string {
	if strings.ToLower(input[1]) == "x" {
		return []string{strings.ToLower(input[0]), "a"}
	} else if strings.ToLower(input[1]) == "y" {
		return []string{strings.ToLower(input[0]), "b"}
	} else {
		return []string{strings.ToLower(input[0]), "c"}
	}
}
