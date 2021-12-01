package internal

import (
	"errors"
	"strconv"
)

type DayOne struct{}

func (day *DayOne) SolveStarOne(input []string) (string, error) {
	values, err := stringSliceToIntSlice(input)
	if err != nil {
		return "", err
	}

	increasedCounter, err := day.aggregateIncreasedValues(values)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(increasedCounter), nil
}

func (day *DayOne) SolveStarTwo(input []string) (string, error) {
	values, err := stringSliceToIntSlice(input)
	if err != nil {
		return "", err
	}

	values, err = day.combineSlidingWindow(values)

	increasedCounter, err := day.aggregateIncreasedValues(values)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(increasedCounter), nil
}

func (day DayOne) aggregateIncreasedValues(input []int) (int, error) {
	increasedCounter := 0
	buffer := -1

	for _, value := range input {
		if buffer >= 0 && value > buffer {
			increasedCounter++
		}

		buffer = value
	}

	return increasedCounter, nil
}

func (day *DayOne) combineSlidingWindow(input []int) ([]int, error) {
	if len(input) < 3 {
		return nil, errors.New("input array is to short, must be at least of length 3")
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

	return result, nil
}
