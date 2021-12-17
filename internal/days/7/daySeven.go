package seven

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type fuelConsumptionRate int

const (
	linear fuelConsumptionRate = iota
	growing
)

type Solver struct {
	values []int
	fuel   int
	sum    int
	length int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).calculateLowestFuelCost(linear).result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).calculateLowestFuelCost(growing).result()
}

func (solver *Solver) parseInput(input []string) *Solver {
	if len(input) != 1 {
		panic("invalid input lines")
	}

	values := strings.Split(strings.TrimSpace(input[0]), ",")
	solver.values = make([]int, len(values))
	solver.length = len(values)
	for idx, v := range values {
		value, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("invalid input '%s'", v))
		}
		solver.sum += value
		solver.values[idx] = value
	}

	return solver
}

func (solver *Solver) calculateLowestFuelCost(rate fuelConsumptionRate) *Solver {
	position := 0
	cost := math.MaxInt32

	for {
		var newCost int
		if rate == growing {
			newCost = solver.calculateCostGrowing(position)
		} else {
			newCost = solver.calculateCostLinear(position)
		}

		if newCost > cost {
			solver.fuel = cost
			break
		}
		cost = newCost
		position++
	}

	return solver
}

func (solver *Solver) calculateCostGrowing(position int) int {
	sum := 0
	for _, v := range solver.values {
		if position > v {
			for i := 1; i <= position-v; i++ {
				sum += i
			}
		} else if position < v {
			for i := 1; i <= v-position; i++ {
				sum += i
			}
		}
	}
	return sum
}

func (solver *Solver) calculateCostLinear(position int) int {
	sum := 0
	for _, v := range solver.values {
		if position > v {
			sum += position - v
		} else if position < v {
			sum += v - position
		}
	}
	return sum
}

func (solver *Solver) result() string {
	return fmt.Sprintf("%d", solver.fuel)
}
