package solver

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Kharonus/aoc/internal/common"
)

type elf struct {
	inventory []int
}

type caloryContainer struct {
	maxValues []int
}

type Solver struct {
	elves []*elf
}

func (day *Solver) SolveStarOne(input []string) string {
	calories := day.parseInput(input).caloriesOfHeaviest(1)
	return strconv.Itoa(calories)
}

func (day *Solver) SolveStarTwo(input []string) string {
	calories := day.parseInput(input).caloriesOfHeaviest(3)
	return strconv.Itoa(calories)
}

func (day *Solver) parseInput(input []string) *Solver {
	var currentElf = &elf{}

	for _, calory := range input {
		if calory == "" {
			day.elves = append(day.elves, currentElf)
			currentElf = &elf{}
			continue
		}

		value, err := strconv.ParseInt(calory, 10, 32)
		if err != nil {
			panic(fmt.Sprintf("'%s' is not a valid calory value.", calory))
		}
		currentElf.inventory = append(currentElf.inventory, int(value))
	}

	return day
}

func (day *Solver) caloriesOfHeaviest(count int) int {
	var container = common.Reduce(day.elves, func(calories caloryContainer, elf *elf) caloryContainer {
		if elf.calories() > calories.maxValues[0] {
			calories.maxValues[0] = elf.calories()
		}

		sort.Ints(calories.maxValues)
		return calories
	}, caloryContainer{maxValues: make([]int, count)})

	return common.Reduce(container.maxValues, common.Sum, 0)
}

func (e *elf) calories() int {
	return common.Reduce(e.inventory, common.Sum, 0)
}
