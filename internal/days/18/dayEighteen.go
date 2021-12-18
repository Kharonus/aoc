package eighteen

import (
	"fmt"
	"math"
	"strconv"
)

type pair struct {
	number int
	depth  int
	left   *pair
	right  *pair
}

type Solver struct {
	pair         *pair
	summands     []*pair
	maxMagnitude int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).resultMagnitude()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).resultMaxMagnitude()
}

func (solver *Solver) parseInput(input []string) *Solver {
	solver.summands = make([]*pair, 0, len(input))
	for _, line := range input {
		p := parsePair(line, 0)
		solver.summands = append(solver.summands, cp(p))
		solver.pair = add(solver.pair, p)
		solver.pair.reduce()
	}

	solver.findMaxMagnitude()

	return solver
}

func (solver *Solver) findMaxMagnitude() {
	max := math.MinInt
	for _, x := range solver.summands {
		for _, y := range solver.summands {
			if x == y {
				continue
			}

			left := cp(x)
			right := cp(y)

			p := add(left, right)
			p.reduce()
			mag := p.magnitude()
			if mag > max {
				max = mag
			}
		}
	}

	solver.maxMagnitude = max
}

func (p *pair) reduce() {
	for {
		if p.explode() {
			continue
		}

		if p.split() {
			continue
		}

		break
	}
}

func (p *pair) explode() bool {
	pairStack := []*pair{p}

	var leftLiteral, rightLiteral *pair
	exploded := false
	var left, right int

	for len(pairStack) > 0 {
		current := pairStack[0]
		pairStack = pairStack[1:]

		if exploded && current.isLiteral() {
			rightLiteral = current
			break
		}

		if current.isLiteral() {
			leftLiteral = current
			continue
		}

		if current.depth == 4 && !exploded {
			exploded = true
			left = current.left.number
			right = current.right.number
			current.number = 0
			current.left = nil
			current.right = nil

			continue
		}

		pairStack = append([]*pair{current.left, current.right}, pairStack...)
	}

	if exploded {
		if leftLiteral != nil {
			leftLiteral.number += left
		}
		if rightLiteral != nil {
			rightLiteral.number += right
		}
	}

	return exploded
}

func (p *pair) split() bool {
	pairStack := []*pair{p}

	for len(pairStack) > 0 {
		current := pairStack[0]
		pairStack = pairStack[1:]

		if current.isLiteral() {
			if current.number > 9 {
				half := float64(current.number) * 0.5

				current.number = 0
				current.left = &pair{number: int(math.Floor(half)), depth: current.depth + 1}
				current.right = &pair{number: int(math.Ceil(half)), depth: current.depth + 1}

				return true
			}

			continue
		}

		pairStack = append([]*pair{current.left, current.right}, pairStack...)
	}

	return false

}

func add(a, b *pair) *pair {
	if a == nil && b == nil {
		panic("cannot add nil pairs")
	}

	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	a.increaseDepth()
	b.increaseDepth()

	return &pair{left: a, right: b, depth: 0}
}

func cp(p *pair) *pair {
	if p == nil {
		return nil
	}

	return &pair{
		number: p.number,
		depth:  p.depth,
		left:   cp(p.left),
		right:  cp(p.right),
	}
}

func (p *pair) increaseDepth() {
	p.depth++
	if p.left != nil {
		p.left.increaseDepth()
	}
	if p.right != nil {
		p.right.increaseDepth()
	}
}

func parsePair(str string, depth int) *pair {
	value, err := strconv.Atoi(str)
	if err == nil {
		return &pair{number: value, depth: depth}
	}

	pairWithoutOuterBrackets := str[1 : len(str)-1]
	openBrackets := 0
	splitIndex := 0
	for idx, c := range pairWithoutOuterBrackets {
		if openBrackets > 0 {
			switch c {
			case '[':
				openBrackets++
			case ']':
				openBrackets--
			}
			continue
		}

		switch c {
		case '[':
			openBrackets++
			continue
		case ',':
			splitIndex = idx
			break
		default:
			continue
		}
	}

	return &pair{
		left:  parsePair(pairWithoutOuterBrackets[:splitIndex], depth+1),
		right: parsePair(pairWithoutOuterBrackets[splitIndex+1:], depth+1),
		depth: depth,
	}
}

func printPair(p *pair) string {
	if p.left != nil && p.right != nil {
		left := printPair(p.left)
		right := printPair(p.right)

		return fmt.Sprintf("[%s,%s]", left, right)
	} else {
		return fmt.Sprintf("%d", p.number)
	}
}

func (p *pair) isLiteral() bool {
	return p.left == nil || p.right == nil
}

func (p *pair) magnitude() int {
	var left, right int
	if p.left.isLiteral() {
		left = p.left.number
	} else {
		left = p.left.magnitude()
	}

	if p.right.isLiteral() {
		right = p.right.number
	} else {
		right = p.right.magnitude()
	}

	return left*3 + right*2
}

func (solver *Solver) resultMagnitude() string {
	return fmt.Sprintf("%d", solver.pair.magnitude())
}

func (solver *Solver) resultMaxMagnitude() string {
	return fmt.Sprintf("%d", solver.maxMagnitude)
}
