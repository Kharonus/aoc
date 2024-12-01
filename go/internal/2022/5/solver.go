package _5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Kharonus/aoc/internal/common"
)

type instruction struct {
	count, source, target int
}

type Solver struct {
	stacks       map[int][]rune
	instructions []*instruction
}

func (day *Solver) SolveStarOne(input []string) string {
	return day.parseInput(input).applyInstructions(true).firstCrates()
}

func (day *Solver) SolveStarTwo(input []string) string {
	return day.parseInput(input).applyInstructions(false).firstCrates()
}

func (day *Solver) parseInput(input []string) *Solver {
	day.stacks = make(map[int][]rune)

	var separatorIndex = 0
	for idx, line := range input {
		if line[1] == '1' {
			separatorIndex = idx + 1
			break
		}

		crates := parseOneLevelOfCrates(line)
		for i, r := range crates {
			day.stacks[i] = append(day.stacks[i], r)
		}
	}

	for _, line := range input[separatorIndex+1:] {
		i := parseInstruction(line)
		day.instructions = append(day.instructions, &i)
	}

	return day
}

func (day *Solver) applyInstructions(reverseOrder bool) *Solver {
	for _, i := range day.instructions {
		day.applyInstruction(i, reverseOrder)
	}

	return day
}

func (day *Solver) applyInstruction(i *instruction, reverseOrder bool) *Solver {
	crates := append(make([]rune, 0), day.stacks[i.source][:i.count]...)
	if reverseOrder {
		common.Reverse(crates)
	}

	day.stacks[i.source] = day.stacks[i.source][i.count:]
	day.stacks[i.target] = append(crates, day.stacks[i.target]...)

	return day
}

func (day *Solver) firstCrates() string {
	result := make([]rune, len(day.stacks))

	for idx, runes := range day.stacks {
		result[idx-1] = runes[0]
	}

	return string(result)
}

func parseOneLevelOfCrates(line string) map[int]rune {
	var result = make(map[int]rune)

	r := regexp.MustCompile(`\[[A-Z]]`)

	matches := r.FindAllString(line, -1)
	indices := r.FindAllStringIndex(line, -1)

	for idx, match := range matches {
		stackNumber := indices[idx][0] / 4
		r2 := []rune(match)[1]
		result[stackNumber+1] = r2
	}

	return result
}

func parseInstruction(line string) instruction {
	var message = fmt.Sprintf("'%s' is not a valid instruction.", line)

	fromIndex := strings.Index(line, "from")
	toIndex := strings.Index(line, "to")

	count, err := strconv.ParseInt(line[5:fromIndex-1], 10, 32)
	if err != nil {
		panic(message)
	}

	source, err := strconv.ParseInt(line[fromIndex+5:toIndex-1], 10, 32)
	if err != nil {
		panic(message)
	}

	target, err := strconv.ParseInt(line[toIndex+3:], 10, 32)
	if err != nil {
		panic(message)
	}

	return instruction{
		count:  int(count),
		source: int(source),
		target: int(target),
	}
}
