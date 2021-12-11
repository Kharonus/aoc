package two

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {
	reach int
	depth int
	aim   int
}

var orderFunctions = map[string]func(day *Solver, value int){
	"forward": func(day *Solver, value int) { day.reach += value },
	"up":      func(day *Solver, value int) { day.depth -= value },
	"down":    func(day *Solver, value int) { day.depth += value },
}

var orderFunctionsWithAim = map[string]func(day *Solver, value int){
	"forward": func(day *Solver, value int) {
		day.reach += value
		day.depth += day.aim * value
	},
	"up":   func(day *Solver, value int) { day.aim -= value },
	"down": func(day *Solver, value int) { day.aim += value },
}

func (solver *Solver) SolveStarOne(input []string) string {
	for _, s := range input {
		solver.parseInput(s, false)
	}

	return solver.result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	for _, s := range input {
		solver.parseInput(s, true)
	}

	return solver.result()
}

func (solver *Solver) parseInput(order string, withAim bool) *Solver {
	splits := strings.Split(strings.TrimSpace(order), " ")
	if len(splits) != 2 {
		panic(fmt.Sprintf("invalid order '%s'", order))
	}

	value, err := strconv.Atoi(splits[1])
	if err != nil {
		panic(fmt.Sprintf("invalid order %s", order))
	}

	var orderFn func(day *Solver, value int)
	if withAim {
		orderFn = orderFunctionsWithAim[splits[0]]
	} else {
		orderFn = orderFunctions[splits[0]]
	}

	if orderFn == nil {
		panic(fmt.Sprintf("invalid order %s", order))
	}

	orderFn(solver, value)
	return solver
}

func (solver *Solver) result() string {
	return fmt.Sprintf("%d", solver.reach*solver.depth)
}
