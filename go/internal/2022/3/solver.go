package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Kharonus/aoc/internal/common"
)

type rucksack struct {
	compartment1, compartment2 string
	allItems                   string
}

type Solver struct {
	rucksacks []*rucksack
}

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (day *Solver) SolveStarOne(input []string) string {
	items := day.parseInput(input).findMisplacedItems()
	return strconv.Itoa(calculatePrioritySum(items))
}

func (day *Solver) SolveStarTwo(input []string) string {
	badges := day.parseInput(input).findBadges()
	return strconv.Itoa(calculatePrioritySum(badges))
}

func (day *Solver) parseInput(input []string) *Solver {
	for _, items := range input {
		var middle = len(items) / 2
		r := &rucksack{
			compartment1: items[:middle],
			compartment2: items[middle:],
			allItems:     items,
		}
		day.rucksacks = append(day.rucksacks, r)
	}

	return day
}

func (day *Solver) findMisplacedItems() string {
	return common.Reduce(day.rucksacks, func(result string, r *rucksack) string {
		intersect := common.IntersectStrings(r.compartment1, r.compartment2)

		if len(intersect) != 1 {
			panic(fmt.Sprintf("Compartments '%s' and '%s' do not have exactly one doubled item.",
				r.compartment1, r.compartment2))
		}

		return result + intersect
	}, "")
}

func (day *Solver) findBadges() string {
	var grouped = common.Reduce(day.rucksacks, func(groups [][]*rucksack, r *rucksack) [][]*rucksack {
		var last = common.Last(groups)
		if len(last) < 3 {
			groups[len(groups)-1] = append(last, r)
		} else {
			groups = append(groups, []*rucksack{r})
		}

		return groups
	}, [][]*rucksack{{}})

	return common.Reduce(grouped, func(result string, group []*rucksack) string {
		var intersect = common.IntersectStrings(group[0].allItems, group[1].allItems)
		intersect = common.IntersectStrings(group[2].allItems, intersect)

		if len(intersect) != 1 {
			panic(fmt.Sprintf("Rucksacks '%s', '%s', and '%s' do not have exactly one shared item.",
				group[0].allItems, group[1].allItems, group[2].allItems))
		}

		return result + intersect
	}, "")
}

func calculatePrioritySum(items string) int {
	return common.Reduce([]rune(items), func(sum int, r rune) int {
		return sum + strings.IndexRune(priority, r) + 1
	}, 0)
}
