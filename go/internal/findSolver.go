package internal

import (
	"fmt"

	_2021_1 "github.com/Kharonus/aoc/internal/2021/1"
	_2021_10 "github.com/Kharonus/aoc/internal/2021/10"
	_2021_11 "github.com/Kharonus/aoc/internal/2021/11"
	_2021_12 "github.com/Kharonus/aoc/internal/2021/12"
	_2021_13 "github.com/Kharonus/aoc/internal/2021/13"
	_2021_14 "github.com/Kharonus/aoc/internal/2021/14"
	_2021_15 "github.com/Kharonus/aoc/internal/2021/15"
	_2021_16 "github.com/Kharonus/aoc/internal/2021/16"
	_2021_17 "github.com/Kharonus/aoc/internal/2021/17"
	_2021_18 "github.com/Kharonus/aoc/internal/2021/18"
	_2021_19 "github.com/Kharonus/aoc/internal/2021/19"
	_2021_2 "github.com/Kharonus/aoc/internal/2021/2"
	_2021_20 "github.com/Kharonus/aoc/internal/2021/20"
	_2021_21 "github.com/Kharonus/aoc/internal/2021/21"
	_2021_22 "github.com/Kharonus/aoc/internal/2021/22"
	_2021_23 "github.com/Kharonus/aoc/internal/2021/23"
	_2021_24 "github.com/Kharonus/aoc/internal/2021/24"
	_2021_25 "github.com/Kharonus/aoc/internal/2021/25"
	_2021_3 "github.com/Kharonus/aoc/internal/2021/3"
	_2021_4 "github.com/Kharonus/aoc/internal/2021/4"
	_2021_5 "github.com/Kharonus/aoc/internal/2021/5"
	_2021_6 "github.com/Kharonus/aoc/internal/2021/6"
	_2021_7 "github.com/Kharonus/aoc/internal/2021/7"
	_2021_8 "github.com/Kharonus/aoc/internal/2021/8"
	_2021_9 "github.com/Kharonus/aoc/internal/2021/9"
	_2022_1 "github.com/Kharonus/aoc/internal/2022/1"
	_2022_10 "github.com/Kharonus/aoc/internal/2022/10"
	_2022_11 "github.com/Kharonus/aoc/internal/2022/11"
	_2022_12 "github.com/Kharonus/aoc/internal/2022/12"
	_2022_13 "github.com/Kharonus/aoc/internal/2022/13"
	_2022_14 "github.com/Kharonus/aoc/internal/2022/14"
	_2022_15 "github.com/Kharonus/aoc/internal/2022/15"
	_2022_16 "github.com/Kharonus/aoc/internal/2022/16"
	_2022_17 "github.com/Kharonus/aoc/internal/2022/17"
	_2022_18 "github.com/Kharonus/aoc/internal/2022/18"
	_2022_19 "github.com/Kharonus/aoc/internal/2022/19"
	_2022_2 "github.com/Kharonus/aoc/internal/2022/2"
	_2022_20 "github.com/Kharonus/aoc/internal/2022/20"
	_2022_21 "github.com/Kharonus/aoc/internal/2022/21"
	_2022_22 "github.com/Kharonus/aoc/internal/2022/22"
	_2022_23 "github.com/Kharonus/aoc/internal/2022/23"
	_2022_24 "github.com/Kharonus/aoc/internal/2022/24"
	_2022_25 "github.com/Kharonus/aoc/internal/2022/25"
	_2022_3 "github.com/Kharonus/aoc/internal/2022/3"
	_2022_4 "github.com/Kharonus/aoc/internal/2022/4"
	_2022_5 "github.com/Kharonus/aoc/internal/2022/5"
	_2022_6 "github.com/Kharonus/aoc/internal/2022/6"
	_2022_7 "github.com/Kharonus/aoc/internal/2022/7"
	_2022_8 "github.com/Kharonus/aoc/internal/2022/8"
	_2022_9 "github.com/Kharonus/aoc/internal/2022/9"
)

type solverCallback func() IDaySolver

var solverMap2021 = map[int]solverCallback{
	1:  func() IDaySolver { return &_2021_1.Solver{} },
	2:  func() IDaySolver { return &_2021_2.Solver{} },
	3:  func() IDaySolver { return &_2021_3.Solver{} },
	4:  func() IDaySolver { return &_2021_4.Solver{} },
	5:  func() IDaySolver { return &_2021_5.Solver{} },
	6:  func() IDaySolver { return &_2021_6.Solver{} },
	7:  func() IDaySolver { return &_2021_7.Solver{} },
	8:  func() IDaySolver { return &_2021_8.Solver{} },
	9:  func() IDaySolver { return &_2021_9.Solver{} },
	10: func() IDaySolver { return &_2021_10.Solver{} },
	11: func() IDaySolver { return &_2021_11.Solver{} },
	12: func() IDaySolver { return &_2021_12.Solver{} },
	13: func() IDaySolver { return &_2021_13.Solver{} },
	14: func() IDaySolver { return &_2021_14.Solver{} },
	15: func() IDaySolver { return &_2021_15.Solver{} },
	16: func() IDaySolver { return &_2021_16.Solver{} },
	17: func() IDaySolver { return &_2021_17.Solver{} },
	18: func() IDaySolver { return &_2021_18.Solver{} },
	19: func() IDaySolver { return &_2021_19.Solver{} },
	20: func() IDaySolver { return &_2021_20.Solver{} },
	21: func() IDaySolver { return &_2021_21.Solver{} },
	22: func() IDaySolver { return &_2021_22.Solver{} },
	23: func() IDaySolver { return &_2021_23.Solver{} },
	24: func() IDaySolver { return &_2021_24.Solver{} },
	25: func() IDaySolver { return &_2021_25.Solver{} },
}

var solverMap2022 = map[int]solverCallback{
	1:  func() IDaySolver { return &_2022_1.Solver{} },
	2:  func() IDaySolver { return &_2022_2.Solver{} },
	3:  func() IDaySolver { return &_2022_3.Solver{} },
	4:  func() IDaySolver { return &_2022_4.Solver{} },
	5:  func() IDaySolver { return &_2022_5.Solver{} },
	6:  func() IDaySolver { return &_2022_6.Solver{} },
	7:  func() IDaySolver { return &_2022_7.Solver{} },
	8:  func() IDaySolver { return &_2022_8.Solver{} },
	9:  func() IDaySolver { return &_2022_9.Solver{} },
	10: func() IDaySolver { return &_2022_10.Solver{} },
	11: func() IDaySolver { return &_2022_11.Solver{} },
	12: func() IDaySolver { return &_2022_12.Solver{} },
	13: func() IDaySolver { return &_2022_13.Solver{} },
	14: func() IDaySolver { return &_2022_14.Solver{} },
	15: func() IDaySolver { return &_2022_15.Solver{} },
	16: func() IDaySolver { return &_2022_16.Solver{} },
	17: func() IDaySolver { return &_2022_17.Solver{} },
	18: func() IDaySolver { return &_2022_18.Solver{} },
	19: func() IDaySolver { return &_2022_19.Solver{} },
	20: func() IDaySolver { return &_2022_20.Solver{} },
	21: func() IDaySolver { return &_2022_21.Solver{} },
	22: func() IDaySolver { return &_2022_22.Solver{} },
	23: func() IDaySolver { return &_2022_23.Solver{} },
	24: func() IDaySolver { return &_2022_24.Solver{} },
	25: func() IDaySolver { return &_2022_25.Solver{} },
}

func FindSolver(year, day int) IDaySolver {
	if day < 1 || day > 25 {
		panic(fmt.Sprintf("Sorry, but day %d is not in my advent calendar", day))
	}

	switch year {
	case 2021:
		return solverMap2021[day]()
	case 2022:
		return solverMap2022[day]()
	default:
		panic(fmt.Sprintf("Sorry, but year %d is not in my advent calendar", year))
	}
}
