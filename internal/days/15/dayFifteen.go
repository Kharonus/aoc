package fifteen

import (
	"fmt"
	"strconv"
)

type pathPoint struct {
	x, y        int
	risk        int
	f           int
	predecessor *pathPoint
}

type Solver struct {
	cave      [][]int
	dimension int
	start     *pathPoint
	end       *pathPoint
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).searchLowRiskPath().result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).enlargeCave().searchLowRiskPath().result()
}

func (solver *Solver) parseInput(input []string) *Solver {
	solver.dimension = len(input)
	solver.cave = make([][]int, len(input))
	for idx, line := range input {
		solver.cave[idx] = make([]int, len(line))
		for charIdx, c := range line {
			value, err := strconv.Atoi(string(c))
			if err != nil {
				panic(fmt.Sprintf("invalid input character '%c' in line '%s'", c, line))
			}

			solver.cave[idx][charIdx] = value
		}
	}

	start := &pathPoint{
		x:           0,
		y:           0,
		risk:        0,
		predecessor: nil,
	}

	start.fScore(solver.hScore(start))
	solver.start = start

	return solver
}

func (solver *Solver) enlargeCave() *Solver {
	initialDimension := solver.dimension
	solver.dimension *= 5

	var initial = make([][]int, initialDimension)
	for i := 0; i < initialDimension; i++ {
		initial[i] = make([]int, initialDimension)
	}

	copy(initial, solver.cave)

	solver.cave = make([][]int, initialDimension*5)
	for i := 0; i < initialDimension*5; i++ {
		solver.cave[i] = make([]int, initialDimension*5)
		for j := 0; j < initialDimension*5; j++ {
			tileY := j / initialDimension
			tileX := i / initialDimension
			offsetY := j % initialDimension
			offsetX := i % initialDimension

			value := initial[offsetX][offsetY] + tileX + tileY

			if value > 9 {
				value -= 9
			}

			solver.cave[i][j] = value
		}
	}

	return solver
}

func (solver *Solver) searchLowRiskPath() *Solver {
	var open = []*pathPoint{solver.start}
	var closed []*pathPoint

	for len(open) > 0 {
		var candidate *pathPoint
		var candidateIdx int
		for idx, p := range open {
			if candidate == nil || candidate.f > p.f {
				candidate = p
				candidateIdx = idx
			}
		}

		open = append(open[:candidateIdx], open[candidateIdx+1:]...)

		successors := solver.getPotentialSuccessors(candidate)
		foundEnd := false
		for _, p := range successors {
			if p.x == solver.dimension-1 && p.y == solver.dimension-1 {
				solver.end = p
				foundEnd = true
				break
			}

			if isInListWithLowerF(open, p) {
				continue
			}

			if isInListWithLowerF(closed, p) {
				continue
			}

			open = append(open, p)
		}

		if foundEnd {
			break
		}

		closed = append(closed, candidate)
	}

	return solver
}

func isInListWithLowerF(list []*pathPoint, p *pathPoint) bool {
	for _, n := range list {
		if n.x == p.x && n.y == p.y && n.f <= p.f {
			return true
		}
	}

	return false
}

func (solver *Solver) getPotentialSuccessors(point *pathPoint) []*pathPoint {
	var successors = make([]*pathPoint, 0, 4)

	if point.x > 0 {
		p := &pathPoint{
			x:           point.x - 1,
			y:           point.y,
			risk:        point.risk + solver.cave[point.x-1][point.y],
			predecessor: point,
		}

		p.fScore(solver.hScore(p))
		successors = append(successors, p)
	}
	if point.x < solver.dimension-1 {

		p := &pathPoint{
			x:           point.x + 1,
			y:           point.y,
			risk:        point.risk + solver.cave[point.x+1][point.y],
			predecessor: point,
		}

		p.fScore(solver.hScore(p))
		successors = append(successors, p)
	}
	if point.y > 0 {
		p := &pathPoint{
			x:           point.x,
			y:           point.y - 1,
			risk:        point.risk + solver.cave[point.x][point.y-1],
			predecessor: point,
		}

		p.fScore(solver.hScore(p))
		successors = append(successors, p)
	}
	if point.y < solver.dimension-1 {
		p := &pathPoint{
			x:           point.x,
			y:           point.y + 1,
			risk:        point.risk + solver.cave[point.x][point.y+1],
			predecessor: point,
		}

		p.fScore(solver.hScore(p))
		successors = append(successors, p)
	}

	return successors
}

func (solver *Solver) hScore(point *pathPoint) int {
	// manhattan distance
	return (solver.dimension - point.x) + (solver.dimension - point.y)
}

func (point *pathPoint) gScore() int {
	return point.risk
}

func (point *pathPoint) fScore(h int) {
	point.f = h + point.gScore()
}

func (solver *Solver) result() string {
	if solver.end == nil {
		return "no path found"
	}

	return fmt.Sprintf("%d", solver.end.risk)
}
