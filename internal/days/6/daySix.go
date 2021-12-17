package six

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {
	swarm [9]int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).reproduce(80).result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).reproduce(256).result()
}

func (solver *Solver) parseInput(input []string) *Solver {
	if len(input) != 1 {
		panic(fmt.Sprintf("invalid input '%+v'", input))
	}

	fishes := strings.Split(strings.TrimSpace(input[0]), ",")
	for _, s := range fishes {
		age, err := strconv.Atoi(s)
		if err != nil || age < 0 || age > 8 {
			panic(fmt.Sprintf("invalid age '%s'", s))
		}

		solver.swarm[age] += 1
	}

	return solver
}

func (solver *Solver) reproduce(days int) *Solver {
	for i := 0; i < days; i++ {
		var newSwarm [9]int
		newSwarm[0] = solver.swarm[1]
		newSwarm[1] = solver.swarm[2]
		newSwarm[2] = solver.swarm[3]
		newSwarm[3] = solver.swarm[4]
		newSwarm[4] = solver.swarm[5]
		newSwarm[5] = solver.swarm[6]
		newSwarm[6] = solver.swarm[7] + solver.swarm[0]
		newSwarm[7] = solver.swarm[8]
		newSwarm[8] = solver.swarm[0]

		solver.swarm = newSwarm
	}

	return solver
}

func (solver *Solver) result() string {
	sum := solver.swarm[0] + solver.swarm[1] + solver.swarm[2] + solver.swarm[3] +
		solver.swarm[4] + solver.swarm[5] + solver.swarm[6] + solver.swarm[7] + solver.swarm[8]
	return fmt.Sprintf("%d", sum)
}
