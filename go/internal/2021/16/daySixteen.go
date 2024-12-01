package sixteen

import (
	"fmt"
	"math"
	"strconv"
)

type Solver struct {
	bits                []int
	accumulatedVersions int
	value               int64
}

type operation int

const (
	sum         operation = 0
	product               = 1
	minimum               = 2
	maximum               = 3
	greaterThan           = 5
	lessThan              = 6
	equalTo               = 7
)

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).iteratePackets().result(int64(solver.accumulatedVersions))
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).iteratePackets().result(solver.value)
}

func (solver *Solver) parseInput(input []string) *Solver {
	if len(input) != 1 {
		panic("invalid input length")
	}

	solver.bits = make([]int, 0, len(input[0])*4)
	for _, c := range input[0] {
		number, err := strconv.ParseInt(string(c), 16, 64)
		if err != nil {
			panic(fmt.Sprintf("invalid hex character: '%c'", c))
		}
		solver.bits = append(solver.bits, int(number>>3&1), int(number>>2&1), int(number>>1&1), int(number&1))
	}

	return solver
}

func (solver *Solver) iteratePackets() *Solver {
	var offset int
	for !solver.remainingBitsAreZeros(offset) {
		length, value := solver.parsePacket(offset)
		offset += length
		solver.value = value
	}

	return solver
}

func (solver *Solver) remainingBitsAreZeros(offset int) bool {
	for i := offset; i < len(solver.bits); i++ {
		if solver.bits[i] == 1 {
			return false
		}
	}

	return true
}

func (solver *Solver) parsePacket(offset int) (length int, value int64) {
	header := solver.bits[offset : offset+6]
	length = 6
	version, typeId := parsePacketHeader(header)
	solver.accumulatedVersions += version

	if typeId == 4 {
		l, v := solver.parseLiteralPacket(offset + 6)
		return length + l, v
	} else {
		l, v := solver.parseOperatorPacket(offset+6, operation(typeId))
		return length + l, v
	}
}

func parsePacketHeader(bits []int) (version, typeId int) {
	version = bitSliceToValue(bits[:3])
	typeId = bitSliceToValue(bits[3:])

	return version, typeId
}

func (solver *Solver) parseLiteralPacket(offset int) (length int, value int64) {
	bitRep := ""
	length = 0

	for i := offset; i < len(solver.bits); i += 5 {
		bitRep += fmt.Sprintf("%d%d%d%d", solver.bits[i+1], solver.bits[i+2], solver.bits[i+3], solver.bits[i+4])
		length += 5
		if solver.bits[i] == 0 {
			break
		}
	}

	v, _ := strconv.ParseInt(bitRep, 2, 64)
	return length, int64(v)
}

func (solver *Solver) parseOperatorPacket(offset int, op operation) (length int, value int64) {
	lengthTypeId := solver.bits[offset]

	length = 0
	var parsedValues []int64

	if lengthTypeId == 0 {
		parsedLength := 0
		totalLength := bitSliceToValue(solver.bits[offset+1 : offset+16])
		for parsedLength < totalLength {
			l, v := solver.parsePacket(offset + 16 + parsedLength)
			parsedLength += l
			parsedValues = append(parsedValues, v)
		}

		length = totalLength + 16
	} else {
		numberOfSubPackets := bitSliceToValue(solver.bits[offset+1 : offset+12])
		totalLength := 12
		for i := 0; i < numberOfSubPackets; i++ {
			l, v := solver.parsePacket(offset + totalLength)
			totalLength += l
			parsedValues = append(parsedValues, v)
		}
		length = totalLength
	}

	return length, calculate(op, parsedValues...)
}

func calculate(op operation, values ...int64) int64 {
	var result int64 = 0

	switch op {
	case sum:
		for _, v := range values {
			result += v
		}
		return result
	case product:
		result = 1
		for _, v := range values {
			result *= v
		}
		return result
	case minimum:
		result = math.MaxInt
		for _, v := range values {
			if v < result {
				result = v
			}
		}
		return result
	case maximum:
		result = math.MinInt
		for _, v := range values {
			if v > result {
				result = v
			}
		}
		return result
	case greaterThan:
		if values[0] > values[1] {
			return 1
		} else {
			return 0
		}
	case lessThan:
		if values[0] < values[1] {
			return 1
		} else {
			return 0
		}
	case equalTo:
		if values[0] == values[1] {
			return 1
		} else {
			return 0
		}
	default:
		panic(fmt.Sprintf("invalid operation '%d'", op))
	}
}

func bitSliceToValue(bits []int) int {
	value := 0
	for idx, bit := range bits {
		value += bit << (len(bits) - 1 - idx)
	}
	return value
}

func (solver *Solver) result(n int64) string {
	return fmt.Sprintf("%d", n)
}
