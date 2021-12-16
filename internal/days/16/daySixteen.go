package sixteen

import (
	"fmt"
	"strconv"
)

type Solver struct {
	bits                []int
	accumulatedVersions int
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).iteratePackets().result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return "What, are you impatient? We do not even approached this far in December 2021 ..."
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
		length := solver.parsePacket(offset)
		offset += length
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

func (solver *Solver) parsePacket(offset int) int {
	packetLength := 0

	header := solver.bits[offset : offset+6]
	packetLength += 6
	version, typeId := parsePacketHeader(header)
	solver.accumulatedVersions += version

	if typeId == 4 {
		_, length := solver.parseLiteralPacket(offset + 6)
		packetLength += length
	} else {
		length := solver.parseOperatorPacket(offset + 6)
		packetLength += length
	}

	return packetLength
}

func parsePacketHeader(bits []int) (version, typeId int) {
	version = bitSliceToValue(bits[:3])
	typeId = bitSliceToValue(bits[3:])

	return version, typeId
}

func (solver *Solver) parseLiteralPacket(offset int) (int, int) {
	bitRep := ""
	length := 0

	for i := offset; i < len(solver.bits); i += 5 {
		bitRep += fmt.Sprintf("%d%d%d%d", solver.bits[i+1], solver.bits[i+2], solver.bits[i+3], solver.bits[i+4])
		length += 5
		if solver.bits[i] == 0 {
			break
		}
	}

	value, _ := strconv.ParseInt(bitRep, 2, 32)
	return int(value), length
}

func (solver *Solver) parseOperatorPacket(offset int) int {
	lengthTypeId := solver.bits[offset]

	if lengthTypeId == 0 {
		totalLength := bitSliceToValue(solver.bits[offset+1 : offset+16])
		return totalLength + 16
	} else {
		numberOfSubPackets := bitSliceToValue(solver.bits[offset+1 : offset+12])
		totalLength := 0
		for i := 0; i < numberOfSubPackets; i++ {
			totalLength += solver.parsePacket(offset + 12 + totalLength)
		}
		return totalLength + 12
	}
}

func bitSliceToValue(bits []int) int {
	value := 0
	for idx, bit := range bits {
		value += bit << (len(bits) - 1 - idx)
	}
	return value
}

func (solver *Solver) result() string {
	return fmt.Sprintf("%d", solver.accumulatedVersions)
}
