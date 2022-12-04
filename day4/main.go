package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, _ := ReadInput()

	input := []string{}

	for _, line := range strings.Split(data, "\n") {
		input = append(input, line)
	}

	assignmentPairsTotalOverlap := 0
	assignmentPairsSomeOverlap := 0

	for _, group := range input {
		group := strings.Split(group, ",")

		var min []int
		var max []int

		for _, elf := range group {
			tasks := strings.Split(elf, "-")

			low, _ := strconv.Atoi(tasks[0])
			high, _ := strconv.Atoi(tasks[1])
			min = append(min, low)
			max = append(max, high)
		}

		if min[0] <= min[1] && max[0] >= max[1] {
			assignmentPairsTotalOverlap++
		} else if min[1] <= min[0] && max[1] >= max[0] {
			assignmentPairsTotalOverlap++
		}

		if min[0] <= min[1] && min[1] <= max[0] { // if elf 2 minimum is between the range of elf 1
			assignmentPairsSomeOverlap++
			continue
		} else if min[1] <= min[0] && min[0] <= max[1] { // if elf 1 minimum is between the range of elf 2
			assignmentPairsSomeOverlap++
			continue
		} else if max[0] >= max[1] && max[1] >= min[0] { // if elf 2 maximum is between the range of elf 1
			assignmentPairsSomeOverlap++
			continue
		} else if max[1] >= max[0] && max[0] >= min[1] { // if elf 1 maximum is between the range of elf 2
			assignmentPairsSomeOverlap++
			continue
		} else {
		}
	}

	fmt.Println("Part 1: ", assignmentPairsTotalOverlap)
	fmt.Println("Part 2: ", assignmentPairsSomeOverlap)
}
