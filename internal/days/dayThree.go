package days

import (
	"fmt"
	"strconv"
	"strings"
)

type DayThree struct {
	values    []int
	bitCounts []int
}

func (day *DayThree) SolveStarOne(input []string) string {
	return day.parseInput(input).getPowerConsumption()
}

func (day *DayThree) SolveStarTwo(input []string) string {
	return "What, are you impatient? We do not even approached this far in December 2021 ..."
}

func (day *DayThree) parseInput(input []string) *DayThree {
	if len(input) == 0 {
		panic("empty input")
	}

	day.values = make([]int, len(input))
	day.bitCounts = make([]int, len(strings.TrimSpace(input[0])))

	for idx, str := range input {
		value, err := strconv.ParseInt(str, 2, 16)
		if err != nil {
			panic(fmt.Sprintf("invalid input %s", str))
		}
		day.values[idx] = int(value)
		increaseBits(&day.bitCounts, day.values[idx])
	}

	return day
}

func (day *DayThree) getPowerConsumption() string {
	var gamma = 0
	var epsilon = 0
	var threshold = len(day.values) / 2

	for idx, v := range day.bitCounts {
		if v > threshold {
			gamma |= pow2(idx)
		} else {
			epsilon |= pow2(idx)
		}
	}

	return fmt.Sprintf("%d", gamma*epsilon)
}

func pow2(x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result *= 2
	}
	return result
}

func increaseBits(aggregate *[]int, value int) {
	high := 1

	for idx := range *aggregate {
		if (value & high) == high {
			(*aggregate)[idx]++
		}
		high *= 2
	}
}
