package _12

import (
	"strconv"
	"strings"
)

type coord [2]int

type location struct {
	coordinate coord
	elevation  int
	pathTo     []*location
}

type Solver struct {
	area       map[coord]*location
	start, end *location
}

const elevationCode = "abcdefghijklmnopqrstuvwxyz"

func (day *Solver) SolveStarOne(input []string) string {
	steps := day.parseInput(input).findPath()
	return strconv.Itoa(steps)
}

func (day *Solver) SolveStarTwo(input []string) string {
	return "What, are you impatient? We didn't even reach this date yet."
}

func (day *Solver) parseInput(input []string) *Solver {
	day.area = map[coord]*location{}
	for y, line := range input {
		for x, r := range line {
			c := coord{x, y}

			switch r {
			case 'S':
				l := location{coordinate: c, elevation: 0}
				day.area[c] = &l
				day.start = &l
			case 'E':
				l := location{coordinate: c, elevation: 23}
				day.area[c] = &l
				day.end = &l
			default:
				day.area[c] = &location{coordinate: c, elevation: strings.IndexRune(elevationCode, r)}
			}
		}
	}

	for c, p := range day.area {
		var neighbors []*location

		if dst, ok := day.area[coord{c.x() - 1, c.y()}]; ok {
			neighbors = append(neighbors, dst)
		}
		if dst, ok := day.area[coord{c.x() + 1, c.y()}]; ok {
			neighbors = append(neighbors, dst)
		}
		if dst, ok := day.area[coord{c.x(), c.y() - 1}]; ok {
			neighbors = append(neighbors, dst)
		}
		if dst, ok := day.area[coord{c.x(), c.y() + 1}]; ok {
			neighbors = append(neighbors, dst)
		}

		for _, neighbor := range neighbors {
			if neighbor.elevation <= p.elevation+1 {
				p.pathTo = append(p.pathTo, neighbor)
			}
		}
	}

	return day
}

func (day *Solver) findPath() int {
	var steps = 0
	var visited = map[*location]int{
		day.start: 0,
	}

	var coordMap = map[int][]*location{
		0: {day.start},
	}

	for len(visited) < len(day.area) {
		for _, l := range coordMap[steps] {
			for _, loc := range l.pathTo {
				if loc == day.end {
					return steps + 1
				}

				if _, ok := visited[loc]; !ok {
					visited[loc] = steps + 1
					coordMap[steps+1] = append(coordMap[steps+1], loc)
				}
			}
		}

		steps += 1
	}

	panic("Couldn't reach end with a path.")
}

func (c coord) x() int {
	return c[0]
}

func (c coord) y() int {
	return c[1]
}
