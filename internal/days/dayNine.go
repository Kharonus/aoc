package days

import (
	"fmt"
	"sort"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

type DayNine struct {
	field     [][]int
	lowPoints []*coordinate
	basinMap  map[coordinate]int
}

func (day *DayNine) SolveStarOne(input []string) string {
	return day.parseInput(input).findLowPoints().resultRiskLevel()
}

func (day *DayNine) SolveStarTwo(input []string) string {
	return day.parseInput(input).findBasins().resultLargestBasins()
}

func (day *DayNine) parseInput(input []string) *DayNine {
	if len(input) < 1 {
		panic("invalid input size")
	}

	day.field = make([][]int, 0, len(input))

	for idx, line := range input {
		day.field = append(day.field, make([]int, len(line)))

		for i := 0; i < len(line); i++ {
			value, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(fmt.Sprintf("invalid character '%d'", line[i]))
			}

			day.field[idx][i] = value
		}
	}

	return day
}

func (day *DayNine) findLowPoints() *DayNine {
	rightBorderIndex := len(day.field[0]) - 1
	bottomBorderIndex := len(day.field) - 1

	for ri, row := range day.field {
		for ci, value := range row {
			if ci < rightBorderIndex && value >= day.field[ri][ci+1] {
				continue
			} else if ci > 0 && value >= day.field[ri][ci-1] {
				continue
			} else if ri < bottomBorderIndex && value >= day.field[ri+1][ci] {
				continue
			} else if ri > 0 && value >= day.field[ri-1][ci] {
				continue
			}

			day.lowPoints = append(day.lowPoints, &coordinate{x: ri, y: ci})
		}
	}

	return day
}

func (day *DayNine) findBasins() *DayNine {
	day.basinMap = make(map[coordinate]int)
	basinIndex := 1

	for ri, row := range day.field {
		for ci, value := range row {
			if value == 9 {
				continue
			}

			var left, top int

			if ci > 0 {
				left = day.basinMap[coordinate{x: ri, y: ci - 1}]
			}

			if ri > 0 {
				top = day.basinMap[coordinate{x: ri - 1, y: ci}]
			}

			current := coordinate{x: ri, y: ci}
			switch {
			case left+top == 0:
				day.basinMap[current] = basinIndex
				basinIndex++
			case left == 0 && top != 0:
				day.basinMap[current] = top
			case left != 0 && top == 0:
				day.basinMap[current] = left
			case left == top:
				day.basinMap[current] = left
			case left != top:
				day.mergeBasins(left, top, basinIndex)
				day.basinMap[current] = basinIndex
				basinIndex++
			}
		}
	}

	return day
}

func (day *DayNine) mergeBasins(b1, b2, new int) {
	for point, basinIdx := range day.basinMap {
		if basinIdx == b1 || basinIdx == b2 {
			day.basinMap[point] = new
		}
	}
}

func (day *DayNine) resultRiskLevel() string {
	sum := 0

	for _, p := range day.lowPoints {
		sum += day.field[p.x][p.y] + 1
	}

	return fmt.Sprintf("%d", sum)
}

func (day *DayNine) resultLargestBasins() string {
	basinSize := make(map[int]int)
	for _, basinIdx := range day.basinMap {
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
