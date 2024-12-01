package four

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

type Solver struct {
	values       []int
	boards       []*board
	winnerRating int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).applyNumbersUntilFirstSolved().result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).applyNumbersUntilLastSolved().result()
}

func (solver *Solver) parseInput(input []string) *Solver {
	validInputLength := len(input) > 6 && ((len(input)-1)%6) == 0

	if !validInputLength {
		panic("input has invalid length")
	}

	solver.values = parseLine(input[0], ",")

	lineNumber := 1
	for lineNumber < len(input) {
		solver.boards = append(solver.boards, parseBoard(input[lineNumber+1:lineNumber+6]))
		lineNumber += 6
	}

	return solver
}

func (solver *Solver) applyNumbersUntilFirstSolved() *Solver {
	var solved = func() *board {
		for _, b := range solver.boards {
			if b.isSolved() {
				return b
			}
		}
		return nil
	}

	round := 0
	var solvedBoard *board
	for solvedBoard == nil && round < len(solver.values) {
		for _, b := range solver.boards {
			b.markNumber(solver.values[round])
		}
		solvedBoard = solved()
		round++
	}

	solver.winnerRating = solvedBoard.getRating(solver.values[round-1])

	return solver
}

func (solver *Solver) applyNumbersUntilLastSolved() *Solver {
	var unsolved = func() []*board {
		var result []*board
		for _, b := range solver.boards {
			if !b.isSolved() {
				result = append(result, b)
			}
		}
		return result
	}

	round := 0
	var unsolvedBoards = solver.boards
	for len(unsolvedBoards) > 1 && round < len(solver.values) {
		for _, b := range solver.boards {
			b.markNumber(solver.values[round])
		}
		unsolvedBoards = unsolved()
		round++
	}

	lastBoard := unsolvedBoards[0]
	for !lastBoard.isSolved() && round < len(solver.values) {
		for _, b := range solver.boards {
			b.markNumber(solver.values[round])
		}
		round++
	}

	solver.winnerRating = unsolvedBoards[0].getRating(solver.values[round-1])

	return solver
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

func (solver *Solver) result() string {
	return fmt.Sprintf("%d", solver.winnerRating)
}
