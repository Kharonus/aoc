package seventeen

import (
	"fmt"
	"strconv"
	"strings"
)

type trench struct {
	x, y [2]int
}

type Solver struct {
	trench       *trench
	highestReach int
	options      int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).findAllFireOptions().result(solver.highestReach)
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).findAllFireOptions().result(solver.options)
}

func (solver *Solver) parseInput(input []string) *Solver {
	if len(input) != 1 {
		panic("invalid input")
	}

	split := strings.Split(input[0][13:], ", ")
	xValues := strings.Split(split[0][2:], "..")
	yValues := strings.Split(split[1][2:], "..")

	solver.trench = &trench{
		x: [2]int{atoi(xValues[0]), atoi(xValues[1])},
		y: [2]int{atoi(yValues[0]), atoi(yValues[1])},
	}

	return solver
}

func (solver *Solver) findAllFireOptions() *Solver {
	for x := 0; x <= solver.trench.x[1]; x++ {
		for y := solver.trench.y[0]; y < 500; y++ {
			if solver.hitsTrench(x, y) {
				solver.options++

				maxReach := maxHigh(y)
				if maxReach > solver.highestReach {
					solver.highestReach = maxReach
				}
			}
		}
	}

	return solver
}

func (solver *Solver) hitsTrench(vx, vy int) bool {
	isIn := func(x, y int) bool {
		return x >= solver.trench.x[0] && x <= solver.trench.x[1] && y >= solver.trench.y[0] && y <= solver.trench.y[1]
	}

	var px, py int
	vX := vx
	vY := vy

	for !isIn(px, py) {
		if px > solver.trench.x[1] || py < solver.trench.y[0] {
			return false
		}

		px += vX
		py += vY

		if vX > 0 {
			vX--
		}

		vY--
	}

	return true
}

func maxHigh(vy int) int {
	sum := 0
	for i := 1; i <= vy; i++ {
		sum += i
	}

	return sum
}

func (solver *Solver) result(value int) string {
	return fmt.Sprintf("%d", value)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid coordinate value '%s'", s))
	}
	return i
}
