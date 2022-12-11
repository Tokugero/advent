package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// elves[elf[calories], elf[calories], elf[calories]]
	// for elf in elves: sum(elf[calories])
	data, _ := ReadInput()
	elves := generateElves(data)
	sort.Ints(elves)

	var elfCount int = 3
	var topCalories int = sum(elves[len(elves)-(elfCount) : len(elves)])

	fmt.Println(elves[len(elves)-(elfCount) : len(elves)])
	fmt.Println("Most calories: ", elves[len(elves)-1])
	fmt.Println("Top 3: ", topCalories)
}

func generateElves(rawData string) []int {
	var elves []int
	var elf []int

	for _, line := range strings.Split(strings.TrimSuffix(rawData, "\n"), "\n") {
		if line != "" {
			intLine, _ := strconv.Atoi(line)
			elf = append(elf, intLine)
		} else {
			elves = append(elves, sum(elf))
			elf = []int{}
		}
	}

	return elves
}

func sum(array []int) int {
	var result int = 0
	for _, value := range array {
		result += int(value)
	}
	return result
}
