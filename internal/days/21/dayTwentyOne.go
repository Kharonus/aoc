package twentyone

import (
	"fmt"
	"strconv"
)

type player struct {
	score, position int
}

type Solver struct {
	p1, p2    *player
	diceRolls int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).rollUntilWin().result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).rollQuantum().result()
}

func (solver *Solver) rollUntilWin() *Solver {
	activePlayer := solver.p1
	die := 1

	for !solver.p1.hasWon() && !solver.p2.hasWon() {
		sum, face := rollDieThrice(die)
		solver.diceRolls += 3
		die = face

		position := activePlayer.position + sum
		activePlayer.position = (position-1)%10 + 1
		activePlayer.score += activePlayer.position

		if activePlayer == solver.p1 {
			activePlayer = solver.p2
		} else {
			activePlayer = solver.p1
		}
	}

	return solver
}

func (solver *Solver) rollQuantum() *Solver {
	
}

func rollDieThrice(face int) (sum, newFace int) {
	for i := 0; i < 3; i++ {
		sum += face
		if face < 100 {
			face++
		} else {
			face = 1
		}
	}

	return sum, face
}

func (p *player) hasWon() bool {
	return p.score >= 1000
}

func (solver *Solver) parseInput(input []string) *Solver {
	solver.p1 = parsePlayerStart(input[0])
	solver.p2 = parsePlayerStart(input[1])
	return solver
}

func parsePlayerStart(line string) *player {
	pos, err := strconv.Atoi(line[28:])
	if err != nil {
		panic(fmt.Sprintf("invalid player starting position '%s'", line))
	}

	return &player{
		score:    0,
		position: pos,
	}
}

func (solver *Solver) result() string {
	var loser *player
	if solver.p1.score > solver.p2.score {
		loser = solver.p2
	} else {
		loser = solver.p1
	}

	return fmt.Sprintf("%d", loser.score*solver.diceRolls)
}
