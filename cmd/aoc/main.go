package main

import (
	"flag"
	"fmt"
	"github.com/Kharonus/aoc/internal/days/1"
	"github.com/Kharonus/aoc/internal/days/12"
	"github.com/Kharonus/aoc/internal/days/2"
	"github.com/Kharonus/aoc/internal/days/5"
	"log"

	"github.com/Kharonus/aoc/internal"
	"github.com/Kharonus/aoc/internal/days"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Error in execution: %+v", r)
		}
	}()

	var day, star int
	var inputFile string

	flag.IntVar(&day, "d", 0, "Specify the day of the advent of code challenge. Must be between 1 and 25.")
	flag.IntVar(&star, "s", 0, "Specify the star of the advent of code challenge. Must be 1 or 2.")
	flag.StringVar(&inputFile, "i", ".", "The input file of the advent of code challenge.")
	flag.Parse()

	if day < 1 || day > 25 {
		log.Fatalf("%d is not a valid day of the advent of code.", day)
	}

	if star < 1 || star > 2 {
		log.Fatalf("%d is not a valid star of the advent of code.", star)
	}

	input := internal.ReadFileLineByLine(inputFile)

	daySolver := findSolver(day)
	var solution string
	switch star {
	case 1:
		solution = daySolver.SolveStarOne(input)
	case 2:
		solution = daySolver.SolveStarTwo(input)
	}

	fmt.Printf(solution)
}

func findSolver(day int) internal.IDaySolver {
	var daySolver internal.IDaySolver

	switch day {
	case 1:
		daySolver = &one.Solver{}
	case 2:
		daySolver = &two.Solver{}
	case 3:
		daySolver = &days.DayThree{}
	case 4:
		daySolver = &days.DayFour{}
	case 5:
		daySolver = &five.Solver{}
	case 6:
		daySolver = &days.DaySix{}
	case 7:
		daySolver = &days.DaySeven{}
	case 8:
		daySolver = &days.DayEight{}
	case 9:
		daySolver = &days.DayNine{}
	case 10:
		daySolver = &days.DayTen{}
	case 11:
		daySolver = &days.DayEleven{}
	case 12:
		daySolver = &twelve.Solver{}
	case 13:
		daySolver = &days.DayThirteen{}
	case 14:
		daySolver = &days.DayFourteen{}
	case 15:
		daySolver = &days.DayFifteen{}
	case 16:
		daySolver = &days.DaySixteen{}
	case 17:
		daySolver = &days.DaySeventeen{}
	case 18:
		daySolver = &days.DayEighteen{}
	case 19:
		daySolver = &days.DayNineteen{}
	case 20:
		daySolver = &days.DayTwenty{}
	case 21:
		daySolver = &days.DayTwentyOne{}
	case 22:
		daySolver = &days.DayTwentyTwo{}
	case 23:
		daySolver = &days.DayTwentyThree{}
	case 24:
		daySolver = &days.DayTwentyFour{}
	case 25:
		daySolver = &days.DayTwentyFive{}
	default:
		panic(fmt.Sprintf("ouch, there is no day %d in the advent calendar", day))
	}

	return daySolver
}
