package _14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Kharonus/aoc/internal/common"
)

type coord [2]int

type material int

const (
	air material = iota
	rock
	sand
)

type Solver struct {
	cave             map[coord]material
	src              coord
	deepestRockLevel int
}

func (day *Solver) SolveStarOne(input []string) string {
	units := day.parseInput(input).fillWithSand(false)
	return strconv.Itoa(units)
}

func (day *Solver) SolveStarTwo(input []string) string {
	units := day.parseInput(input).fillWithSand(true)
	return strconv.Itoa(units)
}

func (day *Solver) parseInput(input []string) *Solver {
	day.src = coord{500, 0}
	day.cave = map[coord]material{}
	day.deepestRockLevel = -1
	for _, line := range input {
		day.parseLine(line)
	}

	return day
}

func (day *Solver) parseLine(line string) {
	message := fmt.Sprintf("Invalid point in line '%s'.", line)
	rocks := common.Reduce(strings.Split(line, " -> "), func(points []coord, str string) []coord {
		p := strings.Split(str, ",")
		if len(p) != 2 {
			panic(message)
		}

		x, ok := common.ParseIntDecimal(p[0])
		if !ok {
			panic(message)
		}

		y, ok := common.ParseIntDecimal(p[1])
		if !ok {
			panic(message)
		}

		return append(points, coord{x, y})
	}, []coord{})

	var pos = 0
	for pos < len(rocks)-1 {
		day.drawLineOfRocks(rocks[pos], rocks[pos+1])
		pos += 1
	}
}

func (day *Solver) drawLineOfRocks(start, end coord) {
	dx := end[0] - start[0]
	dy := end[1] - start[1]

	if dx != 0 && dy != 0 {
		panic("Cannot draw diagonal lines of rock.")
	}

	length, f := common.Abs(dx + dy)
	var horizontal = dx != 0

	for i := 0; i <= length; i++ {
		var c coord
		if horizontal {
			c = coord{start[0] + i*f, start[1]}
		} else {
			c = coord{start[0], start[1] + i*f}
		}

		day.cave[c] = rock
		if c[1] > day.deepestRockLevel {
			day.deepestRockLevel = c[1]
		}
	}
}

func (day *Solver) materialAt(c coord, hasFloor bool) material {
	if hasFloor && c[1] == day.deepestRockLevel {
		return rock
	}

	if mat, ok := day.cave[c]; ok {
		return mat
	}

	return air
}

func (day *Solver) fillWithSand(hasFloor bool) (units int) {
	if hasFloor {
		day.deepestRockLevel += 2
	}

	var rest = true
	for rest {
		if day.materialAt(day.src, hasFloor) == sand {
			break
		}

		rest = day.pourSandUnit(hasFloor)
		if rest {
			units++
		}
	}

	return units
}

func (day *Solver) pourSandUnit(hasFloor bool) (rest bool) {
	sandUnit := day.src
	rest = false

	for sandUnit[1] < day.deepestRockLevel {
		newPosition := coord{sandUnit[0], sandUnit[1] + 1}
		if day.materialAt(newPosition, hasFloor) == air {
			sandUnit = newPosition
			continue
		}

		newPosition = coord{sandUnit[0] - 1, sandUnit[1] + 1}
		if day.materialAt(newPosition, hasFloor) == air {
			sandUnit = newPosition
			continue
		}

		newPosition = coord{sandUnit[0] + 1, sandUnit[1] + 1}
		if day.materialAt(newPosition, hasFloor) == air {
			sandUnit = newPosition
			continue
		}

		rest = true
		day.cave[sandUnit] = sand
		break
	}

	return rest
}
