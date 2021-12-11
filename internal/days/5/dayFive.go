package five

import (
	"fmt"
	"strconv"
	"strings"
)

type change int

const (
	increase change = iota
	decrease
	unchanged
)

type point struct {
	x int
	y int
}

type Solver struct {
	vents map[point]int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input, false).result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input, true).result()
}

func (solver *Solver) parseInput(input []string, withDiagonals bool) *Solver {
	solver.vents = map[point]int{}
	for _, s := range input {
		line := solver.parseLine(s)
		if !withDiagonals && !solver.isHorizontal(line) && !solver.isVertical(line) {
			continue
		}

		solver.applyLineToVents(line)
	}

	return solver
}

func (solver *Solver) isHorizontal(line [2]*point) bool {
	return line[0].y == line[1].y
}

func (solver *Solver) isVertical(line [2]*point) bool {
	return line[0].x == line[1].x
}

func (solver *Solver) applyLineToVents(line [2]*point) {
	var startX, startY, endX, endY int
	startX = line[0].x
	startY = line[0].y
	endX = line[1].x
	endY = line[1].y

	var changeX, changeY change

	switch {
	case startX < endX:
		changeX = increase
	case startX > endX:
		changeX = decrease
	case startX == endX:
		changeX = unchanged
	}

	switch {
	case startY < endY:
		changeY = increase
	case startY > endY:
		changeY = decrease
	case startY == endY:
		changeY = unchanged
	}


	run := true
	for run {
		solver.vents[point{x: startX, y: startY}] += 1

		run = startX != endX || startY != endY

		switch changeX {
		case increase:
			startX++
		case decrease:
			startX--
		}

		switch changeY {
		case increase:
			startY++
		case decrease:
			startY--
		}
	}
}

func (solver *Solver) result() string {
	count := 0
	for _, value := range solver.vents {
		if value > 1 {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

func (solver *Solver) parseLine(line string) [2]*point {
	points := strings.Split(line, "->")
	if len(points) != 2 {
		panic(fmt.Sprintf("invalid input '%s'", line))
	}

	var result [2]*point
	result[0] = solver.parsePoint(points[0])
	result[1] = solver.parsePoint(points[1])

	return result
}

func (solver *Solver) parsePoint(s string) *point {
	coords := strings.Split(strings.TrimSpace(s), ",")
	if len(coords) != 2 {
		panic(fmt.Sprintf("invalid point input '%s'", s))
	}

	p := point{}
	var err error

	p.x, err = strconv.Atoi(strings.TrimSpace(coords[0]))
	if err != nil {
		panic(fmt.Sprintf("invalid coord input '%s'", coords[0]))
	}

	p.y, err = strconv.Atoi(strings.TrimSpace(coords[1]))
	if err != nil {
		panic(fmt.Sprintf("invalid coord input '%s'", coords[1]))
	}

	return &p
}

func (solver *Solver) String() string {
	result := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			value := solver.vents[point{x: j, y: i}]
			if value == 0 {
				result += "."
			} else {
				result += fmt.Sprintf("%d", value)
			}
		}
		result += "\n"
	}

	return result
}
