package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type race struct {
	time     int
	distance int
	margin   int
}

func main() {
	utils.GetInput()
	input, err := utils.ReadInput()
	if err != nil {
		panic(err)
	}

	races := []race{}
	raceP2 := race{}

	for _, line := range input {

		if line == "" {
			continue
		}
		numbers := strings.Split(line, ": ")[1]
		fullNumber := "" // string to hold trimmed number
		if strings.Contains(line, "Time:") {
			for _, number := range strings.Fields(numbers) {
				time, _ := strconv.Atoi(number)
				races = append(races, []race{{time: time}}...)

				fullNumber += number
				raceP2.time, _ = strconv.Atoi(fullNumber)
			}
		} else if strings.Contains(line, "Distance:") {
			for i, number := range strings.Fields(numbers) {
				races[i].distance, _ = strconv.Atoi(number)

				fullNumber += number
				raceP2.distance, _ = strconv.Atoi(fullNumber)
			}
		}
	}

	part1 := 1
	for _, race := range races {
		for hold := 1; hold < race.time; hold++ { // dont' start with zero time or end with zero distance
			travel := (race.time - hold) * hold

			if travel > race.distance {
				race.margin += 1
			}
		}

		part1 *= race.margin
	}

	// part 2
	for hold := 1; hold < raceP2.time; hold++ { // dont' start with zero time or end with zero distance
		travel := (raceP2.time - hold) * hold

		if travel > raceP2.distance {
			if raceP2.time%2 == 0 {
				raceP2.margin = raceP2.time - (hold*2 - 1) // double and subtract one
			} else {
				raceP2.margin = raceP2.time - (hold * 2) // double
			}
			break
		}
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", raceP2.margin)
}
