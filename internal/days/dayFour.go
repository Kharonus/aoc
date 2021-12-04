package days

import (
	"fmt"
	"strconv"
	"strings"
)

type boardNumber struct {
	value  int
	marked bool
}

type board [5][5]boardNumber

type DayFour struct {
	values       []int
	boards       []*board
	winnerRating int
}

func (day *DayFour) SolveStarOne(input []string) string {
	return day.parseInput(input).applyNumbersUntilSolved().result()
}

func (day *DayFour) SolveStarTwo(input []string) string {
	return "What, are you impatient? We do not even approached this far in December 2021 ..."
}

func (day *DayFour) parseInput(input []string) *DayFour {
	validInputLength := len(input) > 6 && ((len(input)-1)%6) == 0

	if !validInputLength {
		panic("input has invalid length")
	}

	day.values = parseLine(input[0], ",")

	lineNumber := 1
	for lineNumber < len(input) {
		day.boards = append(day.boards, parseBoard(input[lineNumber+1:lineNumber+6]))
		lineNumber += 6
	}

	return day
}

func (day *DayFour) applyNumbersUntilSolved() *DayFour {
	var solved = func() *board {
		for _, b := range day.boards {
			if b.isSolved() {
				return b
			}
		}
		return nil
	}

	round := 0
	var solvedBoard *board
	for solvedBoard == nil && round < len(day.values) {
		for _, b := range day.boards {
			b.markNumber(day.values[round])
		}
		solvedBoard = solved()
		round++
	}

	day.winnerRating = solvedBoard.getRating(day.values[round-1])

	return day
}

func parseBoard(input []string) *board {
	if len(input) != 5 {
		panic(fmt.Sprintf("invalid board input %s", strings.Join(input, "\n")))
	}

	var b board
	b.fromIntSlice([][]int{
		parseLine(input[0], " "),
		parseLine(input[1], " "),
		parseLine(input[2], " "),
		parseLine(input[3], " "),
		parseLine(input[4], " "),
	})

	return &b
}

func (b *board) fromIntSlice(input [][]int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			b[i][j] = boardNumber{value: input[i][j], marked: false}
		}
	}
}

func (b *board) isSolved() bool {
	for i := 0; i < 5; i++ {
		horizontal := b[0][i].marked && b[1][i].marked && b[2][i].marked && b[3][i].marked && b[4][i].marked
		vertical := b[i][0].marked && b[i][1].marked && b[i][2].marked && b[i][3].marked && b[i][4].marked

		if vertical || horizontal {
			return true
		}
	}

	return false
}

func (b *board) markNumber(number int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j].value == number {
				b[i][j].marked = true
			}
		}
	}
}

func (b *board) getRating(lastNumber int) int {
	var result = 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b[i][j].marked {
				result += b[i][j].value
			}
		}
	}

	return result * lastNumber
}

func parseLine(str, sep string) []int {
	strValues := removeEmptyStrings(strings.Split(strings.TrimSpace(str), sep))

	var values = make([]int, len(strValues))
	for idx, s := range strValues {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("invalid bingo value %s", s))
		}

		values[idx] = v
	}

	return values
}

func removeEmptyStrings(slice []string) []string {
	result := make([]string, 0, len(slice))
	for _, v := range slice {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

func (day *DayFour) result() string {
	return fmt.Sprintf("%d", day.winnerRating)
}
