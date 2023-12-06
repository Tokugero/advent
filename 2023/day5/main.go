package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type sdMap struct {
	sourceName string
	destName   string
	dest       int
	source     int
	shift      int
}

type seedMap struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

type route struct {
	seedToSoil   []sdMap
	soilToFert   []sdMap
	fertToWater  []sdMap
	waterToLight []sdMap
	lightToTemp  []sdMap
	tempToHumid  []sdMap
	humidToLoc   []sdMap
}

func main() {
	utils.GetInput()
	input, err := utils.ReadInput()
	if err != nil {
		panic(err)
	}

	seeds := []int{}
	routes := route{}

	for i, line := range input {
		if strings.Contains(line, "seeds:") {
			seedNumbersString := strings.Split(line, ": ")[1]
			seedNumbers := strings.Split(seedNumbersString, " ")
			for _, seedNumber := range seedNumbers {
				seed, _ := strconv.Atoi(seedNumber)
				seeds = append(seeds, seed)
			}
		}
		if line == "" {
			continue
		}
		if strings.Contains(line, "seed-to-soil map:") {
			insertMap(&routes.seedToSoil, &input, i, "seed", "soil")
		}
		if strings.Contains(line, "soil-to-fertilizer map:") {
			insertMap(&routes.soilToFert, &input, i, "soil", "fertilizer")
		}
		if strings.Contains(line, "fertilizer-to-water map:") {
			insertMap(&routes.fertToWater, &input, i, "fertilizer", "water")
		}
		if strings.Contains(line, "water-to-light map:") {
			insertMap(&routes.waterToLight, &input, i, "water", "light")
		}
		if strings.Contains(line, "light-to-temperature map:") {
			insertMap(&routes.lightToTemp, &input, i, "light", "temperature")
		}
		if strings.Contains(line, "temperature-to-humidity map:") {
			insertMap(&routes.tempToHumid, &input, i, "temperature", "humidity")
		}
		if strings.Contains(line, "humidity-to-location map:") {
			insertMap(&routes.humidToLoc, &input, i, "humidity", "location")
		}
	}

	part1 := routes.humidToLoc[0].source

	for _, seedNumber := range seeds {
		soil := shiftRoute(seedNumber, &routes.seedToSoil)
		fert := shiftRoute(soil, &routes.soilToFert)
		water := shiftRoute(fert, &routes.fertToWater)
		light := shiftRoute(water, &routes.waterToLight)
		temp := shiftRoute(light, &routes.lightToTemp)
		humid := shiftRoute(temp, &routes.tempToHumid)
		loc := shiftRoute(humid, &routes.humidToLoc)

		if loc < part1 {
			part1 = loc
		}
	}

	seedsPart2 := [][]int{}

	for i := 0; i < len(seeds); i += 2 {
		seedsPart2 = append(seedsPart2, []int{seeds[i], seeds[i+1]})
	}

	part2 := routes.humidToLoc[0].source

	for _, seedPart2 := range seedsPart2 {
		for i := seedPart2[0]; i < seedPart2[0]+seedPart2[1]; i++ {
			soil := shiftRoute(i, &routes.seedToSoil)
			fert := shiftRoute(soil, &routes.soilToFert)
			water := shiftRoute(fert, &routes.fertToWater)
			light := shiftRoute(water, &routes.waterToLight)
			temp := shiftRoute(light, &routes.lightToTemp)
			humid := shiftRoute(temp, &routes.tempToHumid)
			loc := shiftRoute(humid, &routes.humidToLoc)

			if loc < part2 {
				part2 = loc
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func processMap(input *[]string, i int, sourceName string, destName string) sdMap {
	rawString := strings.Split((*input)[i], " ")
	dest, _ := strconv.Atoi(rawString[0])
	source, _ := strconv.Atoi(rawString[1])
	shift, _ := strconv.Atoi(rawString[2])

	return sdMap{sourceName, destName, dest, source, shift}
}

func insertMap(function *[]sdMap, input *[]string, i int, sourceName string, destName string) {
	for true {
		i++
		if (*input)[i] == "" {
			break
		}

		*function = append(*function, processMap(input, i, sourceName, destName))
	}
}

func shiftRoute(sourceNumber int, route *[]sdMap) int {
	pos := sourceNumber
	for _, sdMap := range *route {
		sourceShift := sdMap.dest - sdMap.source
		if sourceNumber >= sdMap.source && sourceNumber <= sdMap.source+sdMap.shift {
			pos = sourceNumber + sourceShift
			break
		}
	}
	return pos
}
