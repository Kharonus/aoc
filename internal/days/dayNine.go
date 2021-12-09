package days

import (
	"fmt"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

type DayNine struct {
	field     [][]int
	lowPoints []*coordinate
}

func (day *DayNine) SolveStarOne(input []string) string {
	return day.parseInput(input).findLowPoints().resultRiskLevel()
}

func (day *DayNine) SolveStarTwo(input []string) string {
	return "What, are you impatient? We do not even approached this far in December 2021 ..."
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

func (day *DayNine) resultRiskLevel() string {
	sum := 0

	for _, p := range day.lowPoints {
		sum += day.field[p.x][p.y] + 1
	}

	return fmt.Sprintf("%d", sum)
}
