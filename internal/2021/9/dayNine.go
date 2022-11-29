package nine

import (
	"fmt"
	"sort"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

type Solver struct {
	field     [][]int
	lowPoints []*coordinate
	basinMap  map[coordinate]int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).findLowPoints().resultRiskLevel()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).findBasins().resultLargestBasins()
}

func (solver *Solver) parseInput(input []string) *Solver {
	if len(input) < 1 {
		panic("invalid input size")
	}

	solver.field = make([][]int, 0, len(input))

	for idx, line := range input {
		solver.field = append(solver.field, make([]int, len(line)))

		for i := 0; i < len(line); i++ {
			value, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(fmt.Sprintf("invalid character '%d'", line[i]))
			}

			solver.field[idx][i] = value
		}
	}

	return solver
}

func (solver *Solver) findLowPoints() *Solver {
	rightBorderIndex := len(solver.field[0]) - 1
	bottomBorderIndex := len(solver.field) - 1

	for ri, row := range solver.field {
		for ci, value := range row {
			if ci < rightBorderIndex && value >= solver.field[ri][ci+1] {
				continue
			} else if ci > 0 && value >= solver.field[ri][ci-1] {
				continue
			} else if ri < bottomBorderIndex && value >= solver.field[ri+1][ci] {
				continue
			} else if ri > 0 && value >= solver.field[ri-1][ci] {
				continue
			}

			solver.lowPoints = append(solver.lowPoints, &coordinate{x: ri, y: ci})
		}
	}

	return solver
}

func (solver *Solver) findBasins() *Solver {
	solver.basinMap = make(map[coordinate]int)
	basinIndex := 1

	for ri, row := range solver.field {
		for ci, value := range row {
			if value == 9 {
				continue
			}

			var left, top int

			if ci > 0 {
				left = solver.basinMap[coordinate{x: ri, y: ci - 1}]
			}

			if ri > 0 {
				top = solver.basinMap[coordinate{x: ri - 1, y: ci}]
			}

			current := coordinate{x: ri, y: ci}
			switch {
			case left+top == 0:
				solver.basinMap[current] = basinIndex
				basinIndex++
			case left == 0 && top != 0:
				solver.basinMap[current] = top
			case left != 0 && top == 0:
				solver.basinMap[current] = left
			case left == top:
				solver.basinMap[current] = left
			case left != top:
				solver.mergeBasins(left, top, basinIndex)
				solver.basinMap[current] = basinIndex
				basinIndex++
			}
		}
	}

	return solver
}

func (solver *Solver) mergeBasins(b1, b2, new int) {
	for point, basinIdx := range solver.basinMap {
		if basinIdx == b1 || basinIdx == b2 {
			solver.basinMap[point] = new
		}
	}
}

func (solver *Solver) resultRiskLevel() string {
	sum := 0

	for _, p := range solver.lowPoints {
		sum += solver.field[p.x][p.y] + 1
	}

	return fmt.Sprintf("%d", sum)
}

func (solver *Solver) resultLargestBasins() string {
	basinSize := make(map[int]int)
	for _, basinIdx := range solver.basinMap {
		basinSize[basinIdx] += 1
	}

	var biggestSizes = make([]int, 3)
	for _, size := range basinSize {
		if size > biggestSizes[0] {
			biggestSizes[0] = size
		}

		sort.Ints(biggestSizes)
	}

	return fmt.Sprintf("%d", biggestSizes[0]*biggestSizes[1]*biggestSizes[2])
}
