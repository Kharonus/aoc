package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Kharonus/aoc/internal/common"
)

type _range struct {
	min, max int
}

type pair struct {
	range1, range2 _range
}

type Solver struct {
	pairs []*pair
}

func (day *Solver) SolveStarOne(input []string) string {
	return strconv.Itoa(day.parseInput(input).allFullContained())
}

func (day *Solver) SolveStarTwo(input []string) string {
	return strconv.Itoa(day.parseInput(input).allOverlapped())
}

func (day *Solver) parseInput(input []string) *Solver {
	for _, line := range input {
		var message = fmt.Sprintf("'%s' is not a valid pair input.", line)
		var split = strings.Split(line, ",")
		if len(split) != 2 {
			panic(message)
		}

		var parseRange = func(str string) _range {
			var s = strings.Split(str, "-")
			if len(s) != 2 {
				panic(message)
			}

			min, err := strconv.ParseInt(s[0], 10, 32)
			if err != nil {
				panic(message)
			}

			max, err := strconv.ParseInt(s[1], 10, 32)
			if err != nil {
				panic(message)
			}

			return _range{min: int(min), max: int(max)}
		}

		day.pairs = append(day.pairs, &pair{
			range1: parseRange(split[0]),
			range2: parseRange(split[1]),
		})
	}

	return day
}

func (day *Solver) allFullContained() int {
	return common.Reduce(day.pairs, func(sum int, p *pair) int {
		if isFullContained(p.range1, p.range2) {
			return sum + 1
		}
		return sum
	}, 0)
}

func isFullContained(range1, range2 _range) bool {
	var diffMin = range2.min - range1.min
	var diffMax = range1.max - range2.max

	return diffMin*diffMax >= 0
}

func (day *Solver) allOverlapped() int {
	return common.Reduce(day.pairs, func(sum int, p *pair) int {
		if isOverlapping(p.range1, p.range2) {
			return sum + 1
		}
		return sum
	}, 0)
}

func isOverlapping(range1, range2 _range) bool {
	var diffMin = range2.min - range1.max
	var diffMax = range2.max - range1.min

	return diffMin*diffMax <= 0
}
