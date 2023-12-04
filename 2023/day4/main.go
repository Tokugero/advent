package main

import (
	"fmt"
	"strconv"
	"strings"
	utils "utils"
)

type Card struct {
	picked  []int
	winning []int
	matched int
	copies  int
}

func main() {
	utils.GetInput()
	input, err := utils.ReadInput()
	if err != nil {
		panic(err)
	}

	games := []Card{}

	for _, line := range input {
		if line == "" {
			continue
		}

		values := strings.Split(line, ": ")    // Get just the numbers
		split := strings.Split(values[1], "|") // Split the number section to picks and wins
		picks := strings.Fields(split[0])      // grab just the content islands
		wins := strings.Fields(split[1])

		card := Card{}

		for _, pick := range picks {
			number, _ := strconv.Atoi(string(pick))
			card.picked = append(card.picked, number)
		}

		for _, win := range wins {
			number, _ := strconv.Atoi(string(win))
			card.winning = append(card.winning, number)
		}

		games = append(games, card)
	}

	// These searches are technically linear, and it would go faster with binary search but the
	// set is so low and advent of cyber is calling me to go play over there.
	scores := []int{}

	for _, game := range games {
		score := 0
		for _, pick := range game.picked {
			for _, win := range game.winning {
				if pick == win {
					if score == 0 {
						score = 1
					} else {
						score = score * 2
					}
				}
			}
		}
		scores = append(scores, score)
	}

	part1 := utils.Sum(scores)

	for i, game := range games {
		for _, pick := range game.picked {
			for _, win := range game.winning {
				if pick == win {
					game.matched++
				}
			}
		}

		for win := 1; win <= game.matched; win++ {
			for copy := 0; copy <= game.copies; copy++ {
				games[i+win].copies++
			}
		}
	}

	part2 := len(games)

	for _, game := range games {
		part2 += game.copies
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
