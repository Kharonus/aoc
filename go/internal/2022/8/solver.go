package _8

import (
	"fmt"
	"strconv"

	"github.com/Kharonus/aoc/internal/common"
)

type direction int

const (
	top direction = iota
	left
	right
	bottom
)

type Solver struct {
	trees         [][]int
	visibility    map[[2]int]struct{}
	width, height int
}

func (day *Solver) SolveStarOne(input []string) string {
	visibility := day.parseInput(input).checkGlobalVisibility().visibility
	return strconv.Itoa(len(visibility))
}

func (day *Solver) SolveStarTwo(input []string) string {
	score := day.parseInput(input).findHighestScenicScore()
	return strconv.Itoa(score)
}

func (day *Solver) parseInput(input []string) *Solver {
	day.height = len(input)
	day.width = len(input[0])

	for _, line := range input {
		day.trees = append(day.trees, parseCiphers(line))
	}

	day.visibility = map[[2]int]struct{}{}
	return day
}

func parseCiphers(line string) []int {
	var result []int
	message := fmt.Sprintf("Line '%s' has non numerical characters.", line)

	for _, r := range line {
		cipher, err := strconv.ParseInt(string(r), 10, 32)
		if err != nil {
			panic(message)
		}

		result = append(result, int(cipher))
	}

	return result
}

func (day *Solver) checkGlobalVisibility() *Solver {
	day.visibility = mergeVisibilities(day.checkVisibilityFrom(top), day.visibility)
	day.visibility = mergeVisibilities(day.checkVisibilityFrom(bottom), day.visibility)
	day.visibility = mergeVisibilities(day.checkVisibilityFrom(left), day.visibility)
	day.visibility = mergeVisibilities(day.checkVisibilityFrom(right), day.visibility)

	return day
}

func (day *Solver) checkVisibilityFrom(dir direction) map[[2]int]struct{} {
	var visibility = map[[2]int]struct{}{}
	var currentHeight int

	var checkAndInsert = func(vertical, horizontal int) int {
		h := day.trees[vertical][horizontal]
		if h > currentHeight {
			visibility[[2]int{vertical, horizontal}] = struct{}{}
			currentHeight = h
		}

		return currentHeight
	}

	switch dir {
	case top:
		for j := 0; j < day.width; j++ {
			currentHeight = -1

			for i := 0; i < day.height; i++ {
				currentHeight = checkAndInsert(i, j)
			}
		}
	case left:
		for i := 0; i < day.height; i++ {
			currentHeight = -1

			for j := 0; j < day.width; j++ {
				currentHeight = checkAndInsert(i, j)
			}
		}
	case right:
		for i := 0; i < day.height; i++ {
			currentHeight = -1

			for j := day.width - 1; j >= 0; j-- {
				currentHeight = checkAndInsert(i, j)
			}
		}
	case bottom:
		for j := 0; j < day.width; j++ {
			currentHeight = -1

			for i := day.height - 1; i >= 0; i-- {
				currentHeight = checkAndInsert(i, j)
			}
		}
	}

	return visibility
}

func (day *Solver) findHighestScenicScore() int {
	var score int

	for vertical, row := range day.trees {
		for horizontal := range row {
			scenicScore := day.checkScenicScore([2]int{vertical, horizontal})
			score = common.Max(score, scenicScore)
		}
	}

	return score
}

func (day *Solver) checkScenicScore(coords [2]int) int {
	return day.checkIn(top, coords) *
		day.checkIn(left, coords) *
		day.checkIn(right, coords) *
		day.checkIn(bottom, coords)

}

func (day *Solver) checkIn(dir direction, coords [2]int) int {
	height := day.trees[coords[0]][coords[1]]
	var n = 1

	var next = map[direction]func(steps int) (int, bool){
		top: func(steps int) (int, bool) {
			c := coords[0] - steps
			if c < 0 || c > day.height-1 {
				return 0, false
			}
			return day.trees[c][coords[1]], true
		},
		left: func(steps int) (int, bool) {
			c := coords[1] - steps
			if c < 0 || c > day.width-1 {
				return 0, false
			}
			return day.trees[coords[0]][c], true
		},
		bottom: func(steps int) (int, bool) {
			c := coords[0] + steps
			if c < 0 || c > day.height-1 {
				return 0, false
			}
			return day.trees[c][coords[1]], true
		},
		right: func(steps int) (int, bool) {
			c := coords[1] + steps
			if c < 0 || c > day.width-1 {
				return 0, false
			}
			return day.trees[coords[0]][c], true
		},
	}

	for {
		tree, ok := next[dir](n)
		if !ok {
			return n - 1
		}
		if tree >= height {
			return n
		}
		n += 1
	}
}

func mergeVisibilities(src, dst map[[2]int]struct{}) map[[2]int]struct{} {
	for coords, _ := range src {
		if _, ok := dst[coords]; !ok {
			dst[coords] = struct{}{}
		}
	}

	return dst
}
