package three

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {
	values    []int
	bitCounts []int
}

type maskBits string

const (
	common   maskBits = "common"
	uncommon          = "uncommon"
)

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).getPowerConsumptionRating()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).getLifeSupportRating()
}

func (solver *Solver) parseInput(input []string) *Solver {
	if len(input) == 0 {
		panic("empty input")
	}

	solver.values = make([]int, len(input))
	solver.bitCounts = make([]int, len(strings.TrimSpace(input[0])))

	for idx, str := range input {
		value, err := strconv.ParseInt(str, 2, 16)
		if err != nil {
			panic(fmt.Sprintf("invalid input %s", str))
		}
		solver.values[idx] = int(value)
		increaseBits(&solver.bitCounts, solver.values[idx])
	}

	return solver
}

func (solver *Solver) getPowerConsumptionRating() string {
	var gamma = 0
	var epsilon = 0
	var threshold = len(solver.values) / 2

	for idx, v := range solver.bitCounts {
		if v > threshold {
			gamma |= pow2(idx)
		} else {
			epsilon |= pow2(idx)
		}
	}

	return fmt.Sprintf("%d", gamma*epsilon)
}

func (solver *Solver) getLifeSupportRating() string {
	oxygenValues := make([]int, len(solver.values))
	copy(oxygenValues, solver.values)
	high := pow2(len(solver.bitCounts) - 1)

	for len(oxygenValues) > 1 && high >= 1 {
		oxygenValues = filterByMaskBit(oxygenValues, high, common)
		high = high / 2
	}

	carbonValues := make([]int, len(solver.values))
	copy(carbonValues, solver.values)
	high = pow2(len(solver.bitCounts) - 1)

	for len(carbonValues) > 1 && high >= 1 {
		carbonValues = filterByMaskBit(carbonValues, high, uncommon)
		high = high / 2
	}

	return fmt.Sprintf("%d", oxygenValues[0]*carbonValues[0])
}

func filterByMaskBit(values []int, high int, mask maskBits) []int {
	var withOnes = make([]int, 0, len(values))
	var withZeros = make([]int, 0, len(values))

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
