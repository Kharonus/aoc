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

type maskBits string

const (
	common   maskBits = "common"
	uncommon          = "uncommon"
)

func (day *DayThree) SolveStarOne(input []string) string {
	return day.parseInput(input).getPowerConsumptionRating()
}

func (day *DayThree) SolveStarTwo(input []string) string {
	return day.parseInput(input).getLifeSupportRating()
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

func (day *DayThree) getPowerConsumptionRating() string {
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

func (day *DayThree) getLifeSupportRating() string {
	oxygenValues := make([]int, len(day.values))
	copy(oxygenValues, day.values)
	high := pow2(len(day.bitCounts) - 1)

	for len(oxygenValues) > 1 && high >= 1 {
		oxygenValues = filterByMaskBit(oxygenValues, high, common)
		high = high / 2
	}

	carbonValues := make([]int, len(day.values))
	copy(carbonValues, day.values)
	high = pow2(len(day.bitCounts) - 1)

	for len(carbonValues) > 1 && high >= 1 {
		carbonValues = filterByMaskBit(carbonValues, high, uncommon)
		high = high / 2
	}

	return fmt.Sprintf("%d", oxygenValues[0]*carbonValues[0])
}

func filterByMaskBit(values []int, high int, mask maskBits) []int {
	var withOnes []int
	var withZeros []int

	for _, value := range values {
		if (value & high) == high {
			withOnes = append(withOnes, value)
		} else {
			withZeros = append(withZeros, value)
		}
	}

	if mask == common && len(withOnes) >= len(withZeros) ||
		mask == uncommon && len(withOnes) < len(withZeros) {
		return withOnes
	}

	return withZeros
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
