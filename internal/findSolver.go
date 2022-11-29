package internal

import (
	"fmt"
	one "github.com/Kharonus/aoc/internal/days/1"
	ten "github.com/Kharonus/aoc/internal/days/10"
	eleven "github.com/Kharonus/aoc/internal/days/11"
	twelve "github.com/Kharonus/aoc/internal/days/12"
	thirteen "github.com/Kharonus/aoc/internal/days/13"
	fourteen "github.com/Kharonus/aoc/internal/days/14"
	fifteen "github.com/Kharonus/aoc/internal/days/15"
	sixteen "github.com/Kharonus/aoc/internal/days/16"
	seventeen "github.com/Kharonus/aoc/internal/days/17"
	eighteen "github.com/Kharonus/aoc/internal/days/18"
	nineteen "github.com/Kharonus/aoc/internal/days/19"
	two "github.com/Kharonus/aoc/internal/days/2"
	twenty "github.com/Kharonus/aoc/internal/days/20"
	twentyone "github.com/Kharonus/aoc/internal/days/21"
	twentytwo "github.com/Kharonus/aoc/internal/days/22"
	twentythree "github.com/Kharonus/aoc/internal/days/23"
	twentyfour "github.com/Kharonus/aoc/internal/days/24"
	twentyfive "github.com/Kharonus/aoc/internal/days/25"
	three "github.com/Kharonus/aoc/internal/days/3"
	four "github.com/Kharonus/aoc/internal/days/4"
	five "github.com/Kharonus/aoc/internal/days/5"
	six "github.com/Kharonus/aoc/internal/days/6"
	seven "github.com/Kharonus/aoc/internal/days/7"
	eight "github.com/Kharonus/aoc/internal/days/8"
	nine "github.com/Kharonus/aoc/internal/days/9"
)

type solverCallback func() IDaySolver

var solverMap = map[int]solverCallback{
	1:  func() IDaySolver { return &one.Solver{} },
	2:  func() IDaySolver { return &two.Solver{} },
	3:  func() IDaySolver { return &three.Solver{} },
	4:  func() IDaySolver { return &four.Solver{} },
	5:  func() IDaySolver { return &five.Solver{} },
	6:  func() IDaySolver { return &six.Solver{} },
	7:  func() IDaySolver { return &seven.Solver{} },
	8:  func() IDaySolver { return &eight.Solver{} },
	9:  func() IDaySolver { return &nine.Solver{} },
	10: func() IDaySolver { return &ten.Solver{} },
	11: func() IDaySolver { return &eleven.Solver{} },
	12: func() IDaySolver { return &twelve.Solver{} },
	13: func() IDaySolver { return &thirteen.Solver{} },
	14: func() IDaySolver { return &fourteen.Solver{} },
	15: func() IDaySolver { return &fifteen.Solver{} },
	16: func() IDaySolver { return &sixteen.Solver{} },
	17: func() IDaySolver { return &seventeen.Solver{} },
	18: func() IDaySolver { return &eighteen.Solver{} },
	19: func() IDaySolver { return &nineteen.Solver{} },
	20: func() IDaySolver { return &twenty.Solver{} },
	21: func() IDaySolver { return &twentyone.Solver{} },
	22: func() IDaySolver { return &twentytwo.Solver{} },
	23: func() IDaySolver { return &twentythree.Solver{} },
	24: func() IDaySolver { return &twentyfour.Solver{} },
	25: func() IDaySolver { return &twentyfive.Solver{} },
}

func FindSolverForDay(day int) IDaySolver {
	if day < 1 || day > 25 {
		panic(fmt.Sprintf("Sorry, but day %d is not in my advent calendar", day))
	}

	return solverMap[day]()
}
