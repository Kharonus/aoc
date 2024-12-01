package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Kharonus/aoc/internal"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Error in execution: %+v", r)
		}
	}()

	var year, day, star int
	var inputFile string

	flag.IntVar(&year, "y", 0, "Specify the year of the advent of code challenge. Must be between 2021 and 2022.")
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

	daySolver := internal.FindSolver(year, day)
	var solution string
	switch star {
	case 1:
		solution = daySolver.SolveStarOne(input)
	case 2:
		solution = daySolver.SolveStarTwo(input)
	}

	fmt.Println(solution)
}
