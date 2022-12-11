package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Register struct {
	register     int
	runningTotal []int
}

func (r *Register) add(registerRaw string) {
	register, _ := strconv.Atoi(registerRaw)
	r.runningTotal = append(r.runningTotal, register) // increase running total by register
}

func (r *Register) noop() {
	return
}

func (r *Register) postCycle() {
	r.register += r.runningTotal[0]
	r.runningTotal = r.runningTotal[1:]
}

func main() {
	data, _ := ReadInput()

	// before the first cycle
	cpu := Register{1, []int{}}
	total := 0

	for _, line := range data {
		parsed := strings.Split(line, " ")

		//build the queue
		cpu.runningTotal = append(cpu.runningTotal, 0)
		switch parsed[0] {
		case "addx":
			cpu.add(parsed[1])
		case "noop":
			cpu.noop()
		}

	}

	for index := range cpu.runningTotal {
		//beginning of cycle
		//fmt.Println("Start: ", cpu.register, cpu.runningTotal)
		//output
		if index == 19 || (index-19)%40 == 0 {
			fmt.Println(index+1, "|", cpu.register, "|", cpu.register*(index+1), "|", cpu.runningTotal)
			total += cpu.register * (index + 1)
		}

		//after the cycle
		cpu.postCycle()
		//fmt.Println("End: ", cpu.register, cpu.runningTotal)
	}
}
