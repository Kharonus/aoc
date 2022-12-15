package _10

import (
	"fmt"
	"github.com/Kharonus/aoc/internal/common"
	"strconv"
	"strings"
)

type Solver struct {
	cycle, register           int
	accumulatedSignalStrength int
	pixels                    []bool
}

func (day *Solver) SolveStarOne(input []string) string {
	signal := day.run(input).accumulatedSignalStrength
	return strconv.Itoa(signal)
}

func (day *Solver) SolveStarTwo(input []string) string {
	return day.run(input).draw()
}

func (day *Solver) run(input []string) *Solver {
	day.cycle = 0
	day.register = 1
	for _, cmd := range input {
		isNoop := cmd == "noop"

		day.runCycle()
		if !isNoop {
			day.runCycle()
			day.setRegister(cmd)
		}
	}

	return day
}

func (day *Solver) setRegister(cmd string) {
	message := fmt.Sprintf("Invalid command '%s'.", cmd)
	split := strings.Split(cmd, " ")
	if len(split) != 2 {
		panic(message)
	}

	val, ok := common.ParseIntDecimal(split[1])
	if !ok {
		panic(message)
	}

	day.register += val
}

func (day *Solver) propagateSignal() bool {
	switch {
	case day.cycle == 20:
		return true
	case (day.cycle-20)%40 == 0:
		return true
	}

	return false
}

func (day *Solver) runCycle() {
	day.cycle += 1

	if day.propagateSignal() {
		day.accumulatedSignalStrength += day.register * day.cycle
	}

	pos := (day.cycle - 1) % 40
	day.pixels = append(day.pixels, common.Diff(pos, day.register) < 2)
}

func (day *Solver) draw() string {
	result := ""

	for idx, pixel := range day.pixels {
		if idx%40 == 0 {
			result += "\n"
		}

		if pixel {
			result += "#"
		} else {
			result += "."
		}
	}

	return result
}
