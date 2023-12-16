package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type Hand struct {
	cards string
	bid   int
	label string
	score Score
}

type Score struct {
	fiveOfAKind  bool
	fourOfAKind  bool
	fullHouse    bool
	threeOfAKind bool
	twoPair      bool
	pair         bool
}

func main() {
	utils.GetInput()
	input, _ := utils.ReadInput()

	points := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	points2 := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 1, // these are jokers now, and worth less in sort order
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	hands := []Hand{}

	for _, line := range input {
		if line == "" {
			break
		}
		h := Hand{}
		h.cards = strings.Split(line, " ")[0]
		h.bid, _ = strconv.Atoi(strings.Split(line, " ")[1])
		hands = append(hands, h)
	}

	gameSplitsP1 := breakoutSplits(hands, false)
	gameSplitsP2 := breakoutSplits(hands, true)

	// sort hands in gameSplit by point order
	for _, gameSplit := range gameSplitsP1 {
		gameSplit = sortGames(gameSplit, points)
	}

	for _, gameSplit := range gameSplitsP2 {
		gameSplit = sortGames(gameSplit, points2)
	}

	part1Raw := sortSplits(gameSplitsP1)
	part2Raw := sortSplits(gameSplitsP2)

	var part1 int
	var part2 int

	for i, hand := range part1Raw {
		part1 += (len(part1Raw) - i) * hand.bid
	}
	for i, hand := range part2Raw {
		part2 += (len(part2Raw) - i) * hand.bid
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}

func countPairs(hand string, wild bool) Score {
	score := Score{false, false, false, false, false, false}

	handSorted := strings.Split(hand, "")
	sort.Strings(handSorted)
	hand = strings.Join(handSorted, "")

	cardsChecked := ""

	wilds := 0 // default wild cards regardless of need
	if wild {
		// ridiculous edge case, probably more scalable way to find this
		if hand == "JJJJJ" {
			return Score{true, false, false, false, false, false} // we got 5 wilds
		}

		wilds = strings.Count(hand, "J")                 // if we're using wilds, let's get a count in the hand
		hand = strings.ReplaceAll(hand, string("J"), "") // remove any jokers now, they'll count to the best hand in the filter
	}

	for _, card := range hand {
		if strings.Count(hand, string(card))+wilds == 5 {
			score.fiveOfAKind = true
			continue
		}
		if strings.Count(hand, string(card))+wilds == 4 {
			score.fourOfAKind = true
			continue
		}
		if strings.Count(hand, string(card))+wilds == 3 {
			fhCheck := strings.ReplaceAll(hand, string(card), "")
			if countPairs(fhCheck, false).pair {
				score.fullHouse = true
				continue
			}
			score.threeOfAKind = true
			continue
		}
		if strings.Count(hand, string(card))+wilds == 2 {
			fhCheck := strings.ReplaceAll(hand, string(card), "")
			if countPairs(fhCheck, false).threeOfAKind {
				score.fullHouse = true
				continue
			} else if countPairs(fhCheck, false).pair {
				score.twoPair = true
				continue
			}
			score.pair = true
			continue
		}
		cardsChecked += string(card)
	}

	return score
}

func sortGames(games []Hand, points map[string]int) []Hand {
	for i := 0; i < len(games); i++ {
		// thank god for copilot
		for j := i + 1; j < len(games); j++ {
			// TODO: Insert a recursive lookup instead of this garbage
			// Check if first string is higher
			if points[string(games[i].cards[0])] < points[string(games[j].cards[0])] {
				games[i], games[j] = games[j], games[i]
			} else if points[string(games[i].cards[0])] == points[string(games[j].cards[0])] {
				// if the first character is the same in each game, try again on the second character
				if points[string(games[i].cards[1])] < points[string(games[j].cards[1])] {
					games[i], games[j] = games[j], games[i]
				} else if points[string(games[i].cards[1])] == points[string(games[j].cards[1])] {
					// keep going all the way to 5, let's not play 100 card stud, thanks.
					if points[string(games[i].cards[2])] < points[string(games[j].cards[2])] {
						games[i], games[j] = games[j], games[i]
					} else if points[string(games[i].cards[2])] == points[string(games[j].cards[2])] {
						if points[string(games[i].cards[3])] < points[string(games[j].cards[3])] {
							games[i], games[j] = games[j], games[i]
						} else if points[string(games[i].cards[3])] == points[string(games[j].cards[3])] {
							if points[string(games[i].cards[4])] < points[string(games[j].cards[4])] {
								games[i], games[j] = games[j], games[i]
							} else if points[string(games[i].cards[4])] == points[string(games[j].cards[4])] {
								if points[string(games[i].cards[5])] < points[string(games[j].cards[5])] {
									games[i], games[j] = games[j], games[i]
								}
							}
						}
					}
				}
			}
		}
	}
	return games
}

func sortSplits(gameSplits map[string][]Hand) []Hand {
	results := []Hand{}
	for _, gameSplit := range gameSplits["fiveOfAKind"] {
		results = append(results, gameSplit)
	}
	for _, gameSplit := range gameSplits["fourOfAKind"] {
		results = append(results, gameSplit)
	}
	for _, gameSplit := range gameSplits["fullHouse"] {
		results = append(results, gameSplit)
	}
	for _, gameSplit := range gameSplits["threeOfAKind"] {
		results = append(results, gameSplit)
	}
	for _, gameSplit := range gameSplits["twoPair"] {
		results = append(results, gameSplit)
	}
	for _, gameSplit := range gameSplits["pair"] {
		results = append(results, gameSplit)
	}
	for _, gameSplit := range gameSplits["highCard"] {
		results = append(results, gameSplit)
	}

	return results
}

func breakoutSplits(hands []Hand, wild bool) map[string][]Hand {
	results := map[string][]Hand{}
	for _, hand := range hands {
		hand.score = countPairs(hand.cards, wild)
		if hand.score.fiveOfAKind {
			results["fiveOfAKind"] = append(results["fiveOfAKind"], hand)
		} else if hand.score.fourOfAKind {
			results["fourOfAKind"] = append(results["fourOfAKind"], hand)
		} else if hand.score.fullHouse {
			results["fullHouse"] = append(results["fullHouse"], hand)
		} else if hand.score.threeOfAKind {
			results["threeOfAKind"] = append(results["threeOfAKind"], hand)
		} else if hand.score.twoPair {
			results["twoPair"] = append(results["twoPair"], hand)
		} else if hand.score.pair {
			results["pair"] = append(results["pair"], hand)
		} else {
			results["highCard"] = append(results["highCard"], hand)
		}
	}

	return results
}
