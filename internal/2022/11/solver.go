package _11

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/Kharonus/aoc/internal/common"
)

type operator int

const (
	add operator = iota
	multiply
)

type monkey struct {
	items       []int
	operation   func(int) int
	reducer     func(int) int
	test        func(int) int
	divisor     int
	inspections int
}

type Solver struct {
	monkeys []*monkey
	modulo  int
}

func (day *Solver) SolveStarOne(input []string) string {
	business := day.parseInput(input, true).rounds(20).calculateBusiness()
	return strconv.Itoa(business)
}

func (day *Solver) SolveStarTwo(input []string) string {
	business := day.parseInput(input, false).rounds(10000).calculateBusiness()
	return strconv.Itoa(business)
}

func (day *Solver) parseInput(input []string, simple bool) *Solver {
	start := regexp.MustCompile(`Monkey\s[0-9]+:`)

	day.modulo = 1
	for idx, line := range input {
		match := start.MatchString(line)
		if !match {
			continue
		}

		day.monkeys = append(day.monkeys, day.parseMonkey(input[idx+1:idx+6]))
	}

	for _, m := range day.monkeys {
		m.reducer = day.parseReducer(simple)
	}
	return day
}

func (day *Solver) parseMonkey(lines []string) *monkey {
	test, divisor := parseTest(lines[2:5])
	day.modulo *= divisor

	return &monkey{
		items:       parseItems(lines[0]),
		operation:   parseOperatorFunction(lines[1]),
		test:        test,
		divisor:     parseSingleNumber(lines[2]),
		inspections: 0,
	}
}

func (day *Solver) rounds(count int) *Solver {
	for i := 0; i < count; i++ {
		for _, m := range day.monkeys {
			day.turn(m)
		}
	}

	return day
}

func (day *Solver) turn(monkey *monkey) {
	for _, item := range monkey.items {
		monkey.inspections += 1
		newItem := monkey.reducer(monkey.operation(item))
		target := monkey.test(newItem)

		day.monkeys[target].items = append(day.monkeys[target].items, newItem)
	}

	monkey.items = []int{}
}

func (day *Solver) calculateBusiness() int {
	bigBusiness := common.Reduce(day.monkeys, func(max []int, monkey *monkey) []int {
		if monkey.inspections > max[0] {
			max[0] = monkey.inspections
			sort.Ints(max)
		}

		return max
	}, []int{0, 0})

	return common.Reduce(bigBusiness, common.Product, 1)
}

func parseItems(str string) []int {
	itemsStr := regexp.MustCompile(`[0-9]+`).FindAllString(str, -1)
	return common.Reduce(itemsStr, func(list []int, str string) []int {
		if item, ok := common.ParseIntDecimal(str); !ok {
			panic(fmt.Sprintf("Invalid list of items: '%s'", str))
		} else {
			return append(list, item)
		}

	}, []int{})
}

func parseOperatorFunction(line string) func(int) int {
	op, number, isNumber := parseOperator(line)
	switch {
	case op == add && !isNumber:
		return func(i int) int { return i + i }
	case op == multiply && !isNumber:
		return func(i int) int { return i * i }
	case op == add:
		return func(i int) int { return i + number }
	case op == multiply:
		return func(i int) int { return i * number }
	}

	return func(i int) int { return i }
}

func parseOperator(line string) (op operator, number int, isNumber bool) {
	opRegex := regexp.MustCompile(`[+*]`)
	switch opRegex.FindString(line) {
	case "+":
		op = add
	case "*":
		op = multiply
	default:
		panic(fmt.Sprintf("Invalid operation: '%s'", line))
	}

	opIdx := opRegex.FindStringIndex(line)
	number, isNumber = common.ParseIntDecimal(line[opIdx[0]+2:])

	return op, number, isNumber
}

func parseTest(lines []string) (test func(int) int, divisor int) {
	divisor = parseSingleNumber(lines[0])
	monkeyTargetTrue := parseSingleNumber(lines[1])
	monkeyTargetFalse := parseSingleNumber(lines[2])

	return func(i int) int {
		if i%divisor == 0 {
			return monkeyTargetTrue
		} else {
			return monkeyTargetFalse
		}
	}, divisor
}

func parseSingleNumber(line string) int {
	number, _ := common.ParseIntDecimal(regexp.MustCompile(`[0-9]+`).FindString(line))
	return number
}

func (day *Solver) parseReducer(simple bool) func(int) int {
	if simple {
		return func(i int) int { return i / 3 }
	} else {
		return func(i int) int { return i % day.modulo }
	}
}
