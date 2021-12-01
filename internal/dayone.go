package internal

import (
	"fmt"
	"strconv"
)

type DayOne struct{}

func (day *DayOne) SolveStarOne(input []string) string {
	values, err := stringSliceToIntSlice(input)
	if err != nil {
		panic(fmt.Sprintf("input is no list of integer values"))
	}

	increasedCounter := day.aggregateIncreasedValues(values)
	return strconv.Itoa(increasedCounter)
}

func (day *DayOne) SolveStarTwo(input []string) string {
	values, err := stringSliceToIntSlice(input)
	if err != nil {
		panic(fmt.Sprintf("input is no list of integer values"))
	}

	values = day.combineSlidingWindow(values)
	return strconv.Itoa(day.aggregateIncreasedValues(values))
}

func (day DayOne) aggregateIncreasedValues(input []int) int {
	increasedCounter := 0
	buffer := -1

	for _, value := range input {
		if buffer >= 0 && value > buffer {
			increasedCounter++
		}

		buffer = value
	}

	return increasedCounter
}

func (day *DayOne) combineSlidingWindow(input []int) []int {
	if len(input) < 3 {
		panic("input array is to short, must be at least of length 3")
	}

	var result = make([]int, len(input)-2)
	for idx, value := range input {
		if idx < len(input)-2 {
			result[idx] = result[idx] + value
		}

		if idx > 0 && idx < len(input)-1 {
			result[idx-1] = result[idx-1] + value
		}

		if idx > 1 {
			result[idx-2] = result[idx-2] + value
		}
	}

	return result
}
