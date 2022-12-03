package main

import (
	"fmt"
	"strings"
)

func main() {
	data, _ := ReadInput()

	input := []string{}

	for _, line := range strings.Split(data, "\n") {
		input = append(input, line)
	}

	badCount := 0
	badgeCount := 0

	for _, pack := range inventory(input) {
		badCount += incorrect(pack)
	}

	for _, group := range group(input) {
		badgeCount += badges(group)
	}

	fmt.Println(badCount)
	fmt.Println(badgeCount)
}

func inventory(input []string) [][]string {
	inventory := [][]string{}

	for _, rucksack := range input {
		inventory = append(inventory, []string{rucksack[0 : len(rucksack)/2], rucksack[(len(rucksack) / 2):]})
	}

	return inventory
}

func incorrect(inventory []string) int {
	incorrect := 0
	priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	validate := make(map[string]bool)
	checked := []string{}

	for _, packOne := range inventory[0] {
		if validate[string(packOne)] {
			continue
		} else {
			for _, packTwo := range inventory[1] {
				if !validate[string(packTwo)] && packOne == packTwo {
					incorrect += strings.Index(priorities, string(packOne)) + 1
					checked = append(checked, string(packOne))
					validate[string(packOne)] = true
				}
			}
		}
	}

	return incorrect
}

func group(input []string) [][]string {
	groups := [][]string{}
	group := []string{}
	elf := 0

	for _, line := range input {
		group = append(group, line)
		elf++
		if elf == 3 {
			groups = append(groups, group)
			group = []string{}
			elf = 0
			continue
		}
	}

	return groups
}

func badges(group []string) int {
	badges := 0
	priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	validate := make(map[string]bool)
	checked := []string{}

	for _, elfOne := range group[0] {
		if validate[string(elfOne)] {
			continue
		} else {
			for _, elfTwo := range group[1] {
				if !validate[string(elfTwo)] && elfOne == elfTwo {
					checked = append(checked, string(elfOne))
					validate[string(elfTwo)] = true
				}
			}
		}
	}

	for _, elfThree := range group[2] {
		if validate[string(elfThree)] {
			badges += strings.Index(priorities, string(elfThree)) + 1
			break
		}
	}

	return badges
}
