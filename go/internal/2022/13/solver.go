package _13

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Kharonus/aoc/internal/common"
)

type order int

const (
	correct order = iota
	incorrect
	indifferent
)

type pack []interface{}

type pair struct {
	a, b pack
}

type markedPack struct {
	sign int
	p    pack
}

type Solver struct {
	pairs  []*pair
	sorted []markedPack
}

func (day *Solver) SolveStarOne(input []string) string {
	correctPairs := day.parseInput(input).findCorrectPairs()
	return strconv.Itoa(correctPairs)
}

func (day *Solver) SolveStarTwo(input []string) string {
	key := day.parseInput(input).sort().decoderKey()
	return strconv.Itoa(key)
}

func (day *Solver) parseInput(input []string) *Solver {
	var a, b pack = nil, nil

	for _, line := range input {
		if line == "" {
			day.pairs = append(day.pairs, &pair{a: a, b: b})
			a = nil
			b = nil
			continue
		}

		match := parseLine(line)
		if a == nil {
			a = match
		} else {
			b = match
		}
	}

	day.pairs = append(day.pairs, &pair{a: a, b: b})
	return day
}

func parseLine(line string) pack {
	var stack = common.Stack{}
	var number []rune

	var addNumber = func(s *common.Stack) {
		if len(number) == 0 {
			return
		}

		d, _ := common.ParseIntDecimal(string(number))
		number = []rune{}
		elem := append(s.Pop().(pack), d)
		stack.Push(elem)
	}

	for _, r := range line {
		switch r {
		case '[':
			stack.Push(pack{})
		case ']':
			addNumber(&stack)

			elem := stack.Pop()
			if stack.Size() == 0 {
				return elem.(pack)
			}

			top := stack.Pop().(pack)
			stack.Push(append(top, elem))
		case ',':
			addNumber(&stack)
		default:
			number = append(number, r)
		}
	}

	panic(fmt.Sprintf("Last closing bracket not reached."))
}

func (day *Solver) findCorrectPairs() int {
	var correctPairs []int
	for i, p := range day.pairs {
		if compare(p.a, p.b) == correct {
			correctPairs = append(correctPairs, i+1)
		}
	}

	return common.Reduce(correctPairs, common.Sum, 0)
}

func (day *Solver) sort() *Solver {
	for _, p := range day.pairs {
		day.sorted = append(day.sorted, markedPack{sign: 0, p: p.a}, markedPack{sign: 0, p: p.b})
	}
	marker1 := markedPack{sign: 1, p: pack{pack{2}}}
	marker2 := markedPack{sign: 2, p: pack{pack{6}}}
	day.sorted = append(day.sorted, marker1, marker2)

	sort.Slice(day.sorted, func(i, j int) bool {
		o := compare(day.sorted[i].p, day.sorted[j].p)
		return o == correct
	})

	return day
}

func (day *Solver) decoderKey() int {
	var a, b int

	for idx, p := range day.sorted {
		switch p.sign {
		case 1:
			a = idx + 1
		case 2:
			b = idx + 1
		}
	}

	return a * b
}

func compare(a, b interface{}) order {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return compareNumbers(a.(int), b.(int))
		default:
			return compare(pack{a}, b)
		}
	default:
		switch b.(type) {
		case int:
			return compare(a, pack{b})
		}
	}

	// compare two lists
	la := a.(pack)
	lb := b.(pack)

	for len(la) > 0 {
		if len(lb) == 0 {
			return incorrect
		}

		o := compare(la[0], lb[0])
		if o != indifferent {
			return o
		}

		la = la[1:]
		lb = lb[1:]
	}

	if len(lb) > 0 {
		return correct
	}

	return indifferent
}

func compareNumbers(a, b int) order {
	switch {
	case a < b:
		return correct
	case a > b:
		return incorrect
	}

	return indifferent
}
