package _9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Kharonus/aoc/internal/common"
)

type direction int

const (
	up direction = iota
	down
	right
	left
)

type Solver struct {
	instructions []direction
	visited      map[[2]int]struct{}
}

var directionMap = map[string]direction{
	"U": up,
	"D": down,
	"R": right,
	"L": left,
}

func (day *Solver) SolveStarOne(input []string) string {
	visited := day.parseInput(input).moveRope(2).visited
	return strconv.Itoa(len(visited))
}

func (day *Solver) SolveStarTwo(input []string) string {
	visited := day.parseInput(input).moveRope(10).visited
	return strconv.Itoa(len(visited))
}

func (day *Solver) parseInput(input []string) *Solver {
	for _, line := range input {
		message := fmt.Sprintf("Invalid instruction '%s'.", line)
		split := strings.Split(line, " ")
		if len(split) != 2 {
			panic(message)
		}

		n, err := strconv.ParseInt(split[1], 10, 32)
		if err != nil {
			panic(message)
		}

		day.instructions = append(day.instructions, makeInstructions(int(n), directionMap[split[0]])...)
	}

	day.visited = map[[2]int]struct{}{}
	return day
}

func makeInstructions(n int, dir direction) []direction {
	result := make([]direction, n)

	for i := range result {
		result[i] = dir
	}

	return result
}

func (day *Solver) moveRope(knotCount int) *Solver {
	knots := make([][2]int, knotCount)
	for i := range knots {
		knots[i] = [2]int{0, 0}
	}

	for _, instruction := range day.instructions {
		day.visited[knots[knotCount-1]] = struct{}{}

		switch instruction {
		case up:
			knots[0][1] += 1
		case down:
			knots[0][1] -= 1
		case right:
			knots[0][0] += 1
		case left:
			knots[0][0] -= 1
		}

		for i := 0; i < knotCount-1; i++ {
			knots[i+1] = moveNextKnot(knots[i], knots[i+1])
		}
	}

	day.visited[knots[knotCount-1]] = struct{}{}
	return day
}

func moveNextKnot(head, tail [2]int) [2]int {
	d := diff(head, tail)
	var x, y = tail[0], tail[1]

	var dx, dy, fx, fy int
	dx, fx = common.Abs(d[0])
	dy, fy = common.Abs(d[1])

	switch {
	case dx == 2 && dy >= 1:
		x += fx
		y += fy
	case dy == 2 && dx >= 1:
		x += fx
		y += fy
	case dx == 2:
		x += fx
	case dy == 2:
		y += fy
	}

	return [2]int{x, y}
}

func diff(head, tail [2]int) [2]int {
	return [2]int{head[0] - tail[0], head[1] - tail[1]}
}
