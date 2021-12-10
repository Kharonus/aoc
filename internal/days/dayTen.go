package days

import (
	"fmt"
	"sort"
)

type DayTen struct {
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

func (day *DayTen) SolveStarOne(input []string) string {
	return day.checkLines(input).resultSyntaxScore()
}

func (day *DayTen) SolveStarTwo(input []string) string {
	return day.checkLines(input).resultIncompleteLineScore()
}

func (day *DayTen) checkLines(input []string) *DayTen {
	day.incompleteScores = make([]int, 0, len(input))

	for _, line := range input {
		day.checkLine(line)
	}

	return day
}

func (day *DayTen) checkLine(line string) {
	var openRunes = make([]bodyRune, 0, len(line))

	for _, character := range line {
		r := bodyRune(character)
		if day.isOpenRune(r) {
			openRunes = append(openRunes, r)
		} else if day.isCloseRune(r) {
			stackSize := len(openRunes)
			if stackSize == 0 {
				panic(fmt.Sprintf("invalid line '%s'. too early closing character '%c'", line, r))
			}

			last := openRunes[stackSize-1]
			if !day.validBody(last, r) {
				day.score += day.syntaxScore(r)
				return
			}

			openRunes = openRunes[:stackSize-1]
		}
	}

	day.incompleteScores = append(day.incompleteScores, day.incompleteLineScore(openRunes))
}

func (day *DayTen) isOpenRune(r bodyRune) bool {
	return r == openBracket || r == openParenthesis || r == openCurly || r == openDiamond
}

func (day *DayTen) isCloseRune(r bodyRune) bool {
	return r == closeBracket || r == closeParenthesis || r == closeCurly || r == closeDiamond
}

func (day *DayTen) validBody(open, close bodyRune) bool {
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

func (day *DayTen) syntaxScore(r bodyRune) int {
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

func (day *DayTen) incompleteLineScore(openRunes []bodyRune) int {
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

func (day *DayTen) resultSyntaxScore() string {
	return fmt.Sprintf("%d", day.score)
}

func (day *DayTen) resultIncompleteLineScore() string {
	sort.Ints(day.incompleteScores)
	return fmt.Sprintf("%d", day.incompleteScores[len(day.incompleteScores)/2])
}
