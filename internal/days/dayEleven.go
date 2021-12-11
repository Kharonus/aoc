package days

import (
	"fmt"
	"strconv"
)

type DayEleven struct {
	flashes             int
	synchronizationStep int
	octopusField        [10][10]int
}

func (day *DayEleven) SolveStarOne(input []string) string {
	return day.parseInput(input).flashFor(100).resultNumberFlashes()
}

func (day *DayEleven) SolveStarTwo(input []string) string {
	return day.parseInput(input).flashFor(-1).resultSynchronizationStep()
}

func (day *DayEleven) parseInput(input []string) *DayEleven {
	if len(input) != 10 {
		panic("invalid input length")
	}

	for row, line := range input {
		if len(line) != 10 {
			panic(fmt.Sprintf("invalid line length '%s'", line))
		}

		for col, c := range line {
			energy, err := strconv.Atoi(string(c))
			if err != nil {
				panic(fmt.Sprintf("invalid energy level '%c'", c))
			}
			day.octopusField[row][col] = energy
		}
	}

	return day
}

func (day *DayEleven) flashFor(days int) *DayEleven {
	step := 0

	for days == -1 || step < days {
		step++

		for {
			if !day.flashOnce() {
				break
			}
		}

		if days == -1 && day.flashSynchronized() {
			day.synchronizationStep = step
			break
		}

		day.raiseEnergyLevel()
	}

	return day
}

func (day *DayEleven) raiseEnergyLevel() {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			day.octopusField[row][col]++
		}
	}
}

func (day *DayEleven) flashOnce() bool {
	flashed := false

	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if day.octopusField[row][col] >= 9 {
				day.flashOctopus(row, col)
				flashed = true
			}
		}
	}

	return flashed
}

func (day *DayEleven) flashSynchronized() bool {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if day.octopusField[row][col] != -1 {
				return false
			}
		}
	}

	return true
}

func (day *DayEleven) flashOctopus(row, col int) {
	// an octopus with -1 flashed this very round and cannot increase any further until next flash round
	day.octopusField[row][col] = -1
	day.flashes++

	type coordinate struct{ row, col int }
	var neighbors = make([]coordinate, 0, 8)

	if row > 0 && col > 0 {
		neighbors = append(neighbors, coordinate{row: row - 1, col: col - 1})
	}
	if row > 0 && col < 9 {
		neighbors = append(neighbors, coordinate{row: row - 1, col: col + 1})
	}
	if row < 9 && col > 0 {
		neighbors = append(neighbors, coordinate{row: row + 1, col: col - 1})
	}
	if row < 9 && col < 9 {
		neighbors = append(neighbors, coordinate{row: row + 1, col: col + 1})
	}
	if row > 0 {
		neighbors = append(neighbors, coordinate{row: row - 1, col: col})
	}
	if col > 0 {
		neighbors = append(neighbors, coordinate{row: row, col: col - 1})
	}
	if row < 9 {
		neighbors = append(neighbors, coordinate{row: row + 1, col: col})
	}
	if col < 9 {
		neighbors = append(neighbors, coordinate{row: row, col: col + 1})
	}

	for _, n := range neighbors {
		if day.octopusField[n.row][n.col] >= 0 {
			day.octopusField[n.row][n.col] += 1
		}
	}
}

func (day *DayEleven) resultNumberFlashes() string {
	return fmt.Sprintf("%d", day.flashes)
}

func (day *DayEleven) resultSynchronizationStep() string {
	return fmt.Sprintf("%d", day.synchronizationStep)
}
