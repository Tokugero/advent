package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	name           int
	items          []int
	operation      Operation
	test           Operation
	ifs            []int
	iTrue, iFalse  *Monkey
	itemsInspected int
}

type Monkeys map[int]*Monkey

type Operation struct {
	modifier string
	value    int
}

type MonkeySort struct {
	key   string
	value int
}

func (m *Monkey) moveItem(item int, remoteMonkey *Monkey) {
	remoteMonkey.items = append(remoteMonkey.items, item)
	m.items = m.items[1:]
	m.itemsInspected++
}

func main() {
	data, _ := ReadInput()

	fmt.Println("Solve1: ", worry(data, 20, "/", 3))
	fmt.Println("Solve2: ", worry(data, 10_000, "%", 3))

}

func monkeyMath(item int, operation Operation, mod int) int {
	switch operation.modifier {
	case "+":
		return item + operation.value
	case "*":
		if operation.value == 0 { // 0 = non-int parsed, prompt means this should be an old * old value
			return item * item
		} else {
			return item * operation.value
		}
	case "/":
		return item / operation.value
	case "%":
		return item % mod
	}

	return item
}

func worry(data []string, rounds int, mod string, worry int) int {
	monkeys := parseMonkey(data)

	modulo := 1

	for _, monkey := range monkeys {
		modulo *= monkey.test.value
	}

	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			thisMonkey := monkeys[j]
			for len(thisMonkey.items) > 0 {
				// Do inspection & worry calculation
				thisMonkey.items[0] = monkeyMath(thisMonkey.items[0], thisMonkey.operation, 0)
				// Calculate Worry based on base worry divisor
				thisMonkey.items[0] = monkeyMath(thisMonkey.items[0], Operation{mod, worry}, modulo)
				// Pass around items to monkeys
				if thisMonkey.items[0]%thisMonkey.test.value == 0 {
					thisMonkey.moveItem(thisMonkey.items[0], thisMonkey.iTrue)
				} else {
					thisMonkey.moveItem(thisMonkey.items[0], thisMonkey.iFalse)
				}
			}
		}
	}

	var highTouchSort []MonkeySort

	for _, v := range monkeys {
		highTouchSort = append(highTouchSort, MonkeySort{key: strconv.Itoa(v.name), value: v.itemsInspected})
	}

	sort.Slice(highTouchSort, func(i, j int) bool {
		return highTouchSort[i].value > highTouchSort[j].value
	})

	highTouchMultiple := 1

	for _, kv := range highTouchSort[:2] {
		highTouchMultiple *= kv.value
	}

	return highTouchMultiple
}

func parseMonkey(data []string) map[int]*Monkey {
	monkeys := make(Monkeys)

	// Build map of all monkeys, without relationship but with identifiers
	monkeyname := ""
	var parsedMonkeyname int
	for _, line := range data {
		rawEntry := strings.TrimSpace(line)
		entry := strings.Split(rawEntry, " ")
		switch entry[0] {
		case "Monkey":
			monkeyname = entry[1][:len(entry[1])-1]
			parsedMonkeyname, _ = strconv.Atoi(monkeyname)
			monkeys[parsedMonkeyname] = &Monkey{name: parsedMonkeyname, ifs: make([]int, 2)}
		case "Starting":
			for _, item := range entry[2:] {
				if string(item[len(item)-1]) == "," {
					startItem, _ := strconv.Atoi(string(item[:len(item)-1]))
					monkeys[parsedMonkeyname].items = append(monkeys[parsedMonkeyname].items, startItem)
				} else {
					startItem, _ := strconv.Atoi(string(item))
					monkeys[parsedMonkeyname].items = append(monkeys[parsedMonkeyname].items, startItem)
				}
			}
		case "Operation:":
			monkeys[parsedMonkeyname].operation.modifier = entry[4]
			monkeys[parsedMonkeyname].operation.value, _ = strconv.Atoi(entry[5]) // if this fails, it stores 0
		case "Test:":
			monkeys[parsedMonkeyname].test.modifier = "/"
			monkeys[parsedMonkeyname].test.value, _ = strconv.Atoi(entry[3])
		case "If":
			sendMonkeyID, _ := strconv.Atoi(entry[5])
			switch entry[1] {
			case "true:":
				monkeys[parsedMonkeyname].ifs[0] = sendMonkeyID
			case "false:":
				monkeys[parsedMonkeyname].ifs[1] = sendMonkeyID
			}
		default:
			monkeyname = ""
		}
	}

	// After all monkeys exist, build relationship
	for _, monkey := range monkeys {
		monkey.iTrue = monkeys[monkey.ifs[0]]
		monkey.iFalse = monkeys[monkey.ifs[1]]
	}

	return monkeys
}

// 22071328874 -- too low
