package twelve

import (
	"fmt"
	"strings"
)

type caveSize string

const (
	big   caveSize = "big"
	small caveSize = "small"
)

type cave struct {
	name        string
	size        caveSize
	connections []*cave
}

type Solver struct {
	caves []*cave
	paths [][]*cave
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).findAllPaths(false).resultNumberPaths()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).findAllPaths(true).resultNumberPaths()
}

func (solver *Solver) parseInput(input []string) *Solver {
	for _, line := range input {
		solver.parseLine(line)
	}

	return solver
}

func (solver *Solver) parseLine(line string) {
	caves := strings.Split(strings.TrimSpace(line), "-")
	if len(caves) != 2 {
		panic(fmt.Sprintf("invalid line input '%s'", line))
	}

	cave1, cave2 := solver.getOrCreateCaves(caves[0], caves[1])
	cave1.connections = append(cave1.connections, cave2)
	cave2.connections = append(cave2.connections, cave1)
}

func (solver *Solver) getOrCreateCaves(cave1, cave2 string) (*cave, *cave) {
	var c1, c2 *cave
	for _, c := range solver.caves {
		if c.name == cave1 {
			c1 = c
		} else if c.name == cave2 {
			c2 = c
		}
	}

	if c1 == nil {
		c1 = &cave{name: cave1, size: deriveSize(cave1)}
		solver.caves = append(solver.caves, c1)
	}

	if c2 == nil {
		c2 = &cave{name: cave2, size: deriveSize(cave2)}
		solver.caves = append(solver.caves, c2)
	}

	return c1, c2
}

func deriveSize(caveName string) caveSize {
	allUpper := true
	for _, c := range caveName {
		if string(c) == strings.ToLower(string(c)) {
			allUpper = false
			break
		}
	}

	if allUpper {
		return big
	}

	return small
}

func (solver *Solver) findAllPaths(canVisitSingleSmallCaveTwice bool) *Solver {
	var start, end = solver.findStartAndEndCave()
	var unfinishedPaths [][]*cave
	unfinishedPaths = append(unfinishedPaths, []*cave{start})

	for len(unfinishedPaths) != 0 {
		var newUnfinishedPaths = make([][]*cave, 0, len(unfinishedPaths))

		for _, path := range unfinishedPaths {
			lastCave := path[len(path)-1]
			if lastCave == end {
				solver.paths = append(solver.paths, path)
				continue
			}

			for _, connectedCave := range lastCave.connections {
				if canGoThatWay(path, connectedCave, canVisitSingleSmallCaveTwice) {
					var newPath = make([]*cave, len(path))
					copy(newPath, path)
					newUnfinishedPaths = append(newUnfinishedPaths, append(newPath, connectedCave))
				}
			}
		}

		unfinishedPaths = newUnfinishedPaths
	}

	return solver
}

func canGoThatWay(path []*cave, next *cave, canVisitSingleSmallCaveTwice bool) bool {
	if next.size == big {
		return true
	}

	if next.name == "start" {
		return false
	}

	smallCaveVisits := map[*cave]int{}
	anySmallCaveWasVisitedTwice := false

	for _, c := range path {
		if c.size == small {
			smallCaveVisits[c] += 1
		}

		if smallCaveVisits[c] > 1 {
			anySmallCaveWasVisitedTwice = true
		}

		if c == next && !canVisitSingleSmallCaveTwice {
			return false
		}
	}

	if anySmallCaveWasVisitedTwice && smallCaveVisits[next] > 0 {
		return false
	}

	return true
}

func (solver *Solver) findStartAndEndCave() (*cave, *cave) {
	var start, end *cave
	for _, c := range solver.caves {
		if c.name == "start" {
			start = c
		} else if c.name == "end" {
			end = c
		}
	}
	return start, end
}

func (solver *Solver) resultNumberPaths() string {
	return fmt.Sprintf("%d", len(solver.paths))
}
