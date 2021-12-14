package fourteen

import (
	"fmt"
	"math"
	"strings"
)

type Solver struct {
	polymer string
	rules   map[string]string
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).insertFor(10).result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).insertFor(40).result()
}

func (solver *Solver) parseInput(input []string) *Solver {
	if len(input) < 3 {
		panic("invalid input length")
	}

	solver.polymer = input[0]
	solver.rules = make(map[string]string)

	for _, line := range input[2:] {
		solver.parseInsertionRule(line)
	}

	return solver
}

func (solver *Solver) insertFor(steps int) *Solver {
	// iterative
	for step := 0; step < steps; step++ {
		fmt.Println(step)
		solver.expandPolymer()
	}

	// recursive
	//result := ""
	//for i := 0; i < len(solver.polymer)-1; i++ {
	//	polymerExtension := solver.expandRecursive(steps, solver.polymer[i:i+2])
	//	if len(result) == 0 {
	//		result += polymerExtension
	//	} else {
	//		result += polymerExtension[1:]
	//	}
	//}
	//solver.polymer = result

	return solver
}

func (solver *Solver) expandRecursive(steps int, pair string) string {
	if steps == 0 {
		return pair
	}

	insertion := solver.rules[pair]

	left := solver.expandRecursive(steps-1, pair[:1]+insertion)
	right := solver.expandRecursive(steps-1, insertion+pair[1:])

	return left + right[1:]
}

func (solver *Solver) expandPolymer() {
	insertions := make([]string, 0, len(solver.polymer)-1)
	for i := 0; i < len(solver.polymer)-1; i++ {
		insertions = append(insertions, solver.rules[solver.polymer[i:i+2]])
	}

	for i := len(insertions) - 1; i >= 0; i-- {
		solver.polymer = solver.polymer[:i+1] + insertions[i] + solver.polymer[i+1:]
	}
}

func (solver *Solver) parseInsertionRule(line string) {
	split := strings.Split(line, "->")
	if len(split) != 2 {
		panic(fmt.Sprintf("invalid insertion rule input '%s'", line))
	}

	pair := strings.TrimSpace(split[0])
	element := strings.TrimSpace(split[1])
	if len(pair) != 2 || len(element) != 1 {
		panic(fmt.Sprintf("invalid insertion rule input '%s'", line))
	}

	solver.rules[pair] = element
}

func (solver *Solver) result() string {
	elementMap := map[string]int{}

	for _, element := range solver.polymer {
		elementMap[string(element)] += 1
	}

	var max = 0
	var min = math.MaxInt32
	for _, count := range elementMap {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}

	return fmt.Sprintf("%d", max-min)
}
