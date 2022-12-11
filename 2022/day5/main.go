package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := ReadInput()

	containersRaw := []string{}
	actionsRaw := []string{}

	// find split in input
	for index, line := range data {
		if strings.Contains(line, "1   2   3") {
			containersRaw = data[:index]
			actionsRaw = data[index+2:]
			break
		}
	}

	//Part 1
	fmt.Print("Part 1: ")
	for _, result := range moveContainer(cleanMovements(actionsRaw), cleanContainers(containersRaw), true) {
		fmt.Print(string(result[len(result)-1:]))
	}
	fmt.Print("\n")

	//Part 2
	fmt.Print("Part 2: ")
	for _, result := range moveContainer(cleanMovements(actionsRaw), cleanContainers(containersRaw), false) {
		fmt.Print(string(result[len(result)-1:]))
	}
	fmt.Print("\n")
}

func cleanContainers(containersRaw []string) [][]byte {
	containers := [][]byte{}

	for _, container := range containersRaw {
		tempStack := []byte{}
		for index, _ := range container {
			if index%2 == 0 { // check every other character
				if index%4 == 0 { // skip it if it's the 4th character
					continue
				}
				tempStack = append(tempStack, byte(container[index-1])) // this should be the character we want
			}
		}
		containers = append(containers, tempStack)
	}

	sortedContainers := [][]byte{}

	for _, container := range containers {
		for characterIndex, character := range container {
			if len(sortedContainers) <= characterIndex {
				sortedContainers = append(sortedContainers, []byte{})
			}
			if character != 32 {
				sortedContainers[characterIndex] = append([]byte{character}, sortedContainers[characterIndex]...)
			}
		}
	}

	return sortedContainers
}

func cleanMovements(actionsRaw []string) [][]int {
	actions := [][]int{}

	for _, action := range actionsRaw {
		expression := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
		results := expression.FindAllStringSubmatch(action, -1)

		move, _ := strconv.Atoi(results[0][1])
		from, _ := strconv.Atoi(results[0][2])
		to, _ := strconv.Atoi(results[0][3])

		actions = append(actions, []int{move, from, to})

	}

	return actions
}

func moveContainer(actions [][]int, containers [][]byte, c9000 bool) [][]byte {
	for _, action := range actions {
		move := action[0]
		from := action[1] - 1
		to := action[2] - 1

		movingContainer := []byte{}
		// move container
		if c9000 {
			movingContainer = ReverseBytes(containers[from][len(containers[from])-move:])
		} else {
			movingContainer = containers[from][len(containers[from])-move:]
		}
		// from
		containers[from] = containers[from][:len(containers[from])-move]

		// add to
		containers[to] = append(containers[to], movingContainer...)
	}

	return containers
}
