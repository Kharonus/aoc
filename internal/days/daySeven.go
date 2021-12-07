package days

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

type DaySeven struct {
	values []int
	fuel   int
	sum    int
	length int
}

func (day *DaySeven) SolveStarOne(input []string) string {
	return day.parseInput(input).calculateLowestFuelCost(linear).result()
}

func (day *DaySeven) SolveStarTwo(input []string) string {
	return day.parseInput(input).calculateLowestFuelCost(growing).result()
}

func (day *DaySeven) parseInput(input []string) *DaySeven {
	if len(input) != 1 {
		panic("invalid input lines")
	}

	values := strings.Split(strings.TrimSpace(input[0]), ",")
	day.values = make([]int, len(values))
	day.length = len(values)
	for idx, v := range values {
		value, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("invalid input '%s'", v))
		}
		day.sum += value
		day.values[idx] = value
	}

	return day
}

func (day *DaySeven) calculateLowestFuelCost(rate fuelConsumptionRate) *DaySeven {
	position := 0
	cost := math.MaxInt32

	for {
		var newCost int
		if rate == growing {
			newCost = day.calculateCostGrowing(position)
		} else {
			newCost = day.calculateCostLinear(position)
		}

		if newCost > cost {
			day.fuel = cost
			break
		}
		cost = newCost
		position++
	}

	return day
}

func (day *DaySeven) calculateCostGrowing(position int) int {
	sum := 0
	for _, v := range day.values {
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

func (day *DaySeven) calculateCostLinear(position int) int {
	sum := 0
	for _, v := range day.values {
		if position > v {
			sum += position - v
		} else if position < v {
			sum += v - position
		}
	}
	return sum
}

func (day *DaySeven) result() string {
	return fmt.Sprintf("%d", day.fuel)
}
