package solver

import (
	"fmt"
	"github.com/Kharonus/aoc/internal/common"
	"strconv"
	"strings"
)

type shape int8

const (
	rock shape = iota + 1
	paper
	scissor
)

type round struct {
	me, opponent shape
}

type Solver struct {
	rounds []*round
}

var shapeMap = map[string]shape{
	"A": rock,
	"B": paper,
	"C": scissor,
	"X": rock,
	"Y": paper,
	"Z": scissor,
}

func (day *Solver) SolveStarOne(input []string) string {
	var score = day.parseInput(input, false).score()
	return strconv.Itoa(score)
}

func (day *Solver) SolveStarTwo(input []string) string {
	var score = day.parseInput(input, true).score()
	return strconv.Itoa(score)
}

func (day *Solver) parseInput(input []string, deriveShape bool) *Solver {
	for _, r := range input {
		split := strings.Split(r, " ")
		if len(split) != 2 || !strings.Contains("ABC", split[0]) || !strings.Contains("XYZ", split[1]) {
			panic(fmt.Sprintf("'%s' is not a valid round", r))
		}

		currentRound := &round{me: shapeMap[split[1]], opponent: shapeMap[split[0]]}
		if deriveShape {
			currentRound.deriveMyShape(split[1])
		}
		day.rounds = append(day.rounds, currentRound)
	}

	return day
}

func (day *Solver) score() int {
	var sumRounds = func(sum int, r *round) int {
		score := r.score()
		return sum + score
	}
	return common.Reduce(day.rounds, sumRounds, 0)
}

func (r *round) score() int {
	switch {
	case r.me == r.opponent:
		return 3 + int(r.me)
	case r.me == rock && r.opponent == scissor:
		return 6 + int(r.me)
	case r.opponent == rock && r.me == scissor:
		return int(r.me)
	case r.me > r.opponent:
		return 6 + int(r.me)
	default:
		return int(r.me)
	}
}

func (r *round) deriveMyShape(code string) {
	switch code {
	case "Y":
		r.me = r.opponent
	case "X":
		r.me = r.opponent - 1
		if r.me == 0 {
			r.me = scissor
		}
	case "Z":
		r.me = r.opponent + 1
		if r.me == 4 {
			r.me = rock
		}
	}
}
