package main

import (
	"flag"
	"fmt"
	"github.com/Kharonus/aoc/internal/days/20"
	"log"

	"github.com/Kharonus/aoc/internal"
	"github.com/Kharonus/aoc/internal/days"
	"github.com/Kharonus/aoc/internal/days/1"
	"github.com/Kharonus/aoc/internal/days/10"
	"github.com/Kharonus/aoc/internal/days/11"
	"github.com/Kharonus/aoc/internal/days/12"
	"github.com/Kharonus/aoc/internal/days/13"
	"github.com/Kharonus/aoc/internal/days/14"
	"github.com/Kharonus/aoc/internal/days/15"
	"github.com/Kharonus/aoc/internal/days/16"
	"github.com/Kharonus/aoc/internal/days/17"
	"github.com/Kharonus/aoc/internal/days/18"
	"github.com/Kharonus/aoc/internal/days/19"
	"github.com/Kharonus/aoc/internal/days/2"
	"github.com/Kharonus/aoc/internal/days/3"
	"github.com/Kharonus/aoc/internal/days/4"
	"github.com/Kharonus/aoc/internal/days/5"
	"github.com/Kharonus/aoc/internal/days/6"
	"github.com/Kharonus/aoc/internal/days/7"
	"github.com/Kharonus/aoc/internal/days/8"
	"github.com/Kharonus/aoc/internal/days/9"
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
		daySolver = &three.Solver{}
	case 4:
		daySolver = &four.Solver{}
	case 5:
		daySolver = &five.Solver{}
	case 6:
		daySolver = &six.Solver{}
	case 7:
		daySolver = &seven.Solver{}
	case 8:
		daySolver = &eight.Solver{}
	case 9:
		daySolver = &nine.Solver{}
	case 10:
		daySolver = &ten.Solver{}
	case 11:
		daySolver = &eleven.Solver{}
	case 12:
		daySolver = &twelve.Solver{}
	case 13:
		daySolver = &thirteen.Solver{}
	case 14:
		daySolver = &fourteen.Solver{}
	case 15:
		daySolver = &fifteen.Solver{}
	case 16:
		daySolver = &sixteen.Solver{}
	case 17:
		daySolver = &seventeen.Solver{}
	case 18:
		daySolver = &eighteen.Solver{}
	case 19:
		daySolver = &nineteen.Solver{}
	case 20:
		daySolver = &twenty.Solver{}
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
