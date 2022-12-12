package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Register struct {
	register       int
	runningTotal   []int
	signalStrength int
	totalScore     int
}

func (r *Register) add(registerRaw string) {
	register, _ := strconv.Atoi(registerRaw)
	r.noop()
	r.runningTotal = append(r.runningTotal, register) // increase running total by register
}

func (r *Register) noop() {
	r.runningTotal = append(r.runningTotal, 0)
}

func (r *Register) tick() {
	r.register += r.runningTotal[0]
	r.runningTotal = r.runningTotal[1:]
}

func (r *Register) cycle(index int) {
	if index == 19 || (index-19)%40 == 0 {
		r.signalStrength = r.register * (index + 1)
		r.totalScore += r.signalStrength
	}

	r.tick()
}

func main() {
	data, _ := ReadInput()

	// before the first cycle
	cpu := Register{1, []int{}, 0, 0}

	for _, line := range data {
		parsed := strings.Split(line, " ")

		//build the queue
		switch parsed[0] {
		case "addx":
			cpu.add(parsed[1])
		case "noop":
			cpu.noop()
		}
	}

	screen := []bool{}

	for index := range cpu.runningTotal {
		if index < 40 {
			screen = append(screen, (cpu.register == index+1 || cpu.register == index || cpu.register == index-1))
		} else if index < 80 {
			screen = append(screen, (cpu.register == index-40+1 || cpu.register == index-40 || cpu.register == index-40-1))
		} else if index < 120 {
			screen = append(screen, (cpu.register == index-(40*2)+1 || cpu.register == index-(40*2) || cpu.register == index-(40*2)-1))
		} else if index < 160 {
			screen = append(screen, (cpu.register == index-(40*3)+1 || cpu.register == index-(40*3) || cpu.register == index-(40*3)-1))
		} else if index < 200 {
			screen = append(screen, (cpu.register == index-(40*4)+1 || cpu.register == index-(40*4) || cpu.register == index-(40*4)-1))
		} else if index < 240 {
			screen = append(screen, (cpu.register == index-(40*5)+1 || cpu.register == index-(40*5) || cpu.register == index-(40*5)-1))
		}
		cpu.cycle(index)
	}

	for width, pixel := range screen {
		if pixel {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if (width+1)%40 == 0 {
			fmt.Print("\n")
		}
	}

	fmt.Println("Total Score:", cpu.totalScore)
}
