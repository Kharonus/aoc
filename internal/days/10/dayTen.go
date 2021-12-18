package ten

import (
	"fmt"
	"sort"
)

type Solver struct {
	score            int
	incompleteScores []int
}

type bodyRune rune

const (
	openParenthesis  bodyRune = '('
	openBracket      bodyRune = '['
	openCurly        bodyRune = '{'
	openDiamond      bodyRune = '<'
	closeParenthesis bodyRune = ')'
	closeBracket     bodyRune = ']'
	closeCurly       bodyRune = '}'
	closeDiamond     bodyRune = '>'
)

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.checkLines(input).resultSyntaxScore()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.checkLines(input).resultIncompleteLineScore()
}

func (solver *Solver) checkLines(input []string) *Solver {
	solver.incompleteScores = make([]int, 0, len(input))

	for _, line := range input {
		solver.checkLine(line)
	}

	return solver
}

func (solver *Solver) checkLine(line string) {
	var openRunes = make([]bodyRune, 0, len(line))

	for _, character := range line {
		r := bodyRune(character)
		if solver.isOpenRune(r) {
			openRunes = append(openRunes, r)
		} else if solver.isCloseRune(r) {
			stackSize := len(openRunes)
			if stackSize == 0 {
				panic(fmt.Sprintf("invalid line '%s'. too early closing character '%c'", line, r))
			}

			last := openRunes[stackSize-1]
			if !solver.validBody(last, r) {
				solver.score += solver.syntaxScore(r)
				return
			}

			openRunes = openRunes[:stackSize-1]
		}
	}

	solver.incompleteScores = append(solver.incompleteScores, solver.incompleteLineScore(openRunes))
}

func (solver *Solver) isOpenRune(r bodyRune) bool {
	return r == openBracket || r == openParenthesis || r == openCurly || r == openDiamond
}

func (solver *Solver) isCloseRune(r bodyRune) bool {
	return r == closeBracket || r == closeParenthesis || r == closeCurly || r == closeDiamond
}

func (solver *Solver) validBody(open, close bodyRune) bool {
	switch {
	case open == openParenthesis && close == closeParenthesis:
		return true
	case open == openCurly && close == closeCurly:
		return true
	case open == openBracket && close == closeBracket:
		return true
	case open == openDiamond && close == closeDiamond:
		return true
	default:
		return false
	}
}

func (solver *Solver) syntaxScore(r bodyRune) int {
	switch r {
	case closeParenthesis:
		return 3
	case closeBracket:
		return 57
	case closeCurly:
		return 1197
	case closeDiamond:
		return 25137
	default:
		return 0
	}
}

func (solver *Solver) incompleteLineScore(openRunes []bodyRune) int {
	totalScore := 0
	for i := len(openRunes) - 1; i >= 0; i-- {
		totalScore *= 5

		switch openRunes[i] {
		case openParenthesis:
			totalScore += 1
		case openBracket:
			totalScore += 2
		case openCurly:
			totalScore += 3
		case openDiamond:
			totalScore += 4
		}
	}

	return totalScore
}

func (solver *Solver) resultSyntaxScore() string {
	return fmt.Sprintf("%d", solver.score)
}

func (solver *Solver) resultIncompleteLineScore() string {
	sort.Ints(solver.incompleteScores)
	return fmt.Sprintf("%d", solver.incompleteScores[len(solver.incompleteScores)/2])
}
