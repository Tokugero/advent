package main

import (
	"fmt"
	"strings"
	"utils"
)

type Step struct {
	left  string
	right string
}

func main() {
	utils.GetInput()
	input, _ := utils.ReadInput()

	steps := make(map[string]Step)
	var instructions []string

	for i, line := range input {
		if i == 1 {
			continue
		}
		if i == 0 {
			instructions = strings.Split(line, "")
			continue
		}
		if line == "" {
			break
		}
		current := line[0:3]
		left := line[7:10]
		right := line[12:15]

		steps[current] = Step{left, right}
	}

	part1solution := part1("AAA", "ZZZ", steps, instructions)
	part2solution := part2("A", "Z", steps, instructions)

	fmt.Println("Part 1: ", part1solution)
	fmt.Println("Part 2: ", part2solution)
}

func part1(start string, end string, steps map[string]Step, instructions []string) int {
	position := start
	step := 0
	i := 0
	for true {
		if instructions[i] == "L" {
			position = steps[position].left
		} else if instructions[i] == "R" {
			position = steps[position].right
		}

		if position == end {
			step++
			break
		}
		step++
		i++
		if i == len(instructions) {
			i = 0
		}
	}
	return step
}

func part2(start string, end string, steps map[string]Step, instructions []string) int {
	paths := []int{}

	for key := range steps {
		if string(key[2]) == start {
			position := key
			step := 0
			i := 0
			for true {
				if instructions[i] == "L" {
					position = steps[position].left
				} else if instructions[i] == "R" {
					position = steps[position].right
				}

				if string(position[2]) == end {
					step++
					break
				}
				step++
				i++
				if i == len(instructions) {
					i = 0
				}
			}
			paths = append(paths, step)
		}
	}

	return LCM(paths[0], paths[1], paths[2:]...)
}

// TODO: Figure out what this really does. I remembered that path traversal like this can be calculated with
// lowest common multiples and found this snippet online. I understand what it's doing, but need to read more
// on WHY it works.

// credit: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
