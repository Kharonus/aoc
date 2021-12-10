package days

import (
	"fmt"
)

type DayTen struct {
	score int
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
	return day.checkCorruptLines(input).resultSyntaxScore()
}

func (day *DayTen) SolveStarTwo(input []string) string {
	return "What, are you impatient? We do not even approached this far in December 2021 ..."
}

func (day *DayTen) checkCorruptLines(input []string) *DayTen {
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
				break
			}

			openRunes = openRunes[:stackSize-1]
		}
	}
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

func (day *DayTen) resultSyntaxScore() string {
	return fmt.Sprintf("%d", day.score)
}
