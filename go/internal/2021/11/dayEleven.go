package eleven

import (
	"fmt"
	"strconv"
)

type Solver struct {
	flashes             int
	synchronizationStep int
	octopusField        [10][10]int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).flashFor(100).resultNumberFlashes()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).flashFor(-1).resultSynchronizationStep()
}

func (solver *Solver) parseInput(input []string) *Solver {
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
			solver.octopusField[row][col] = energy
		}
	}

	return solver
}

func (solver *Solver) flashFor(days int) *Solver {
	step := 0

	for days == -1 || step < days {
		step++

		for {
			if !solver.flashOnce() {
				break
			}
		}

		if days == -1 && solver.flashSynchronized() {
			solver.synchronizationStep = step
			break
		}

		solver.raiseEnergyLevel()
	}

	return solver
}

func (solver *Solver) raiseEnergyLevel() {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			solver.octopusField[row][col]++
		}
	}
}

func (solver *Solver) flashOnce() bool {
	flashed := false

	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if solver.octopusField[row][col] >= 9 {
				solver.flashOctopus(row, col)
				flashed = true
			}
		}
	}

	return flashed
}

func (solver *Solver) flashSynchronized() bool {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if solver.octopusField[row][col] != -1 {
				return false
			}
		}
	}

	return true
}

func (solver *Solver) flashOctopus(row, col int) {
	// an octopus with -1 flashed this very round and cannot increase any further until next flash round
	solver.octopusField[row][col] = -1
	solver.flashes++

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
		if solver.octopusField[n.row][n.col] >= 0 {
			solver.octopusField[n.row][n.col] += 1
		}
	}
}

func (solver *Solver) resultNumberFlashes() string {
	return fmt.Sprintf("%d", solver.flashes)
}

func (solver *Solver) resultSynchronizationStep() string {
	return fmt.Sprintf("%d", solver.synchronizationStep)
}
