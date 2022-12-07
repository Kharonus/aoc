package _6

import (
	"strconv"
)

type Solver struct{}

func (day *Solver) SolveStarOne(input []string) string {
	return strconv.Itoa(day.findMarker(input[0], 4))
}

func (day *Solver) SolveStarTwo(input []string) string {
	return strconv.Itoa(day.findMarker(input[0], 14))
}

func (day *Solver) findMarker(input string, size int) int {
	queue := []rune(input[:size])

	for idx, r := range input[size:] {
		if !hasDuplicates(queue) {
			return idx + size
		}

		queue = shift(queue, r)
	}

	return 0
}

func hasDuplicates(runes []rune) bool {
	set := map[rune]struct{}{}

	for _, r := range runes {
		if _, ok := set[r]; ok {
			return true
		}

		set[r] = struct{}{}
	}

	return false
}

func shift(runes []rune, r rune) []rune {
	return append(runes[1:], r)
}
