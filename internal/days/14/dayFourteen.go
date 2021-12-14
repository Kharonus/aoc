package fourteen

import (
	"fmt"
	"math"
	"strings"
)

type Solver struct {
	lastElement string
	rules       map[string][]string
	moleculeMap map[string]int
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

	solver.moleculeMap = map[string]int{}
	for i := 0; i < len(input[0])-1; i++ {
		molecule := input[0][i : i+2]
		solver.moleculeMap[molecule] += 1
	}
	solver.lastElement = input[0][len(input[0])-1:]

	solver.rules = make(map[string][]string)
	for _, line := range input[2:] {
		solver.parseRule(line)
	}

	return solver
}

func (solver *Solver) insertFor(steps int) *Solver {
	for step := 0; step < steps; step++ {
		solver.expandPolymer()
	}

	return solver
}

func (solver *Solver) expandPolymer() {
	newMoleculeMap := map[string]int{}

	for molecule, count := range solver.moleculeMap {
		newMolecules := solver.rules[molecule]
		newMoleculeMap[newMolecules[0]] += count
		newMoleculeMap[newMolecules[1]] += count
	}

	solver.moleculeMap = newMoleculeMap
}

func (solver *Solver) parseRule(line string) {
	split := strings.Split(line, "->")
	if len(split) != 2 {
		panic(fmt.Sprintf("invalid insertion rule input '%s'", line))
	}

	pair := strings.TrimSpace(split[0])
	element := strings.TrimSpace(split[1])
	if len(pair) != 2 || len(element) != 1 {
		panic(fmt.Sprintf("invalid insertion rule input '%s'", line))
	}

	solver.rules[pair] = []string{pair[:1] + element, element + pair[1:]}
}

func (solver *Solver) result() string {
	elementMap := map[string]int{}

	for molecule, count := range solver.moleculeMap {
		elementMap[molecule[:1]] += count
	}

	elementMap[solver.lastElement] += 1

	var max = 0
	var min = math.MaxInt64
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
