package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Game struct {
	number int
	rounds []Round
	max    Round
}

type Round struct {
	red   int
	green int
	blue  int
}

func main() {
	utils.GetInput()
	input, _ := utils.ReadInput()

	games := map[int]Game{}

	for _, line := range input {
		if line == "" {
			continue
		}

		game := Game{}

		split := strings.Split(line, ": ")
		gameNumber, _ := strconv.Atoi(strings.Split(split[0], " ")[1])
		game.number = gameNumber

		unparsedRounds := strings.Split(split[1], "; ")
		for _, unparsedRound := range unparsedRounds {
			round := Round{}
			colors := strings.Split(unparsedRound, ", ")

			for _, color := range colors {
				result := strings.Split(color, " ")
				if strings.Contains(result[1], "red") {
					round.red, _ = strconv.Atoi(result[0])
					if round.red > game.max.red {
						game.max.red = round.red
					}
				}
				if strings.Contains(result[1], "green") {
					round.green, _ = strconv.Atoi(result[0])
					if round.green > game.max.green {
						game.max.green = round.green
					}
				}
				if strings.Contains(result[1], "blue") {
					round.blue, _ = strconv.Atoi(result[0])
					if round.blue > game.max.blue {
						game.max.blue = round.blue
					}
				}
			}

			game.rounds = append(game.rounds, round)
		}

		games[gameNumber] = game
	}

	gameInput1 := Round{
		red:   12,
		green: 13,
		blue:  14,
	}

	possibleGames := 0 // Part 1
	powerCubes := 0    // Part 2

	for _, game := range games {
		if (game.max.red <= gameInput1.red) && (game.max.green <= gameInput1.green) && (game.max.blue <= gameInput1.blue) {
			possibleGames += game.number
		}
		powerCubes += game.max.red * game.max.green * game.max.blue
	}

	fmt.Println("Part 1: ", possibleGames)
	fmt.Println("Part 2: ", powerCubes)
}
