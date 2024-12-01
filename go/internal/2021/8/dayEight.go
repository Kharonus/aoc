package eight

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type codedLine struct {
	digits     []string
	segmentMap [7]string
}

type Solver struct {
	digitsUnique []string
	lines        []*codedLine
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).resultUniqueDigits()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).resultDecodedSum()
}

func (solver *Solver) parseInput(input []string) *Solver {
	solver.digitsUnique = make([]string, 0, len(input)*4)
	solver.lines = make([]*codedLine, 0, len(input))
	for _, s := range input {
		segments, digits := solver.parseLine(s)

		solver.lines = append(solver.lines, &codedLine{
			digits:     digits,
			segmentMap: solver.deduceSegments(segments),
		})

		for _, d := range digits {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				solver.digitsUnique = append(solver.digitsUnique, d)
			}
		}
	}

	return solver
}

func (solver *Solver) parseLine(line string) (segments, digits []string) {
	split := strings.Split(strings.TrimSpace(line), "|")
	if len(split) != 2 {
		panic(fmt.Sprintf("invalid line '%s'", line))
	}

	segments = strings.Split(strings.TrimSpace(split[0]), " ")
	digits = strings.Split(strings.TrimSpace(split[1]), " ")
	return
}

func (solver *Solver) resultUniqueDigits() string {
	return fmt.Sprintf("%d", len(solver.digitsUnique))
}

func (solver *Solver) resultDecodedSum() string {
	sum := 0
	for _, line := range solver.lines {
		number := 0
		for i := 0; i < len(line.digits); i++ {
			number += int(math.Pow10(3-i)) * solver.toDigit(line.segmentMap, line.digits[i])
		}
		sum += number
	}

	return fmt.Sprintf("%d", sum)
}

func (solver *Solver) deduceSegments(segments []string) [7]string {
	var segmentMap [7]string
	sorted := solver.sortSegmentsByLength(segments)

	leftSide := solver.uniqueChars(sorted[3:6]...)
	topLeftMiddle := solver.uniqueChars(sorted[0], sorted[2])

	segmentMap[0] = solver.uniqueChars(sorted[0], sorted[1])
	segmentMap[1] = solver.intersect(leftSide, topLeftMiddle)
	segmentMap[3] = solver.uniqueChars(segmentMap[1], topLeftMiddle)
	segmentMap[4] = solver.uniqueChars(leftSide, segmentMap[1])
	segmentMap[6] = solver.uniqueChars(leftSide, segmentMap[3], sorted[1], sorted[9])

	segmentMap[5] = solver.intersect(append(sorted[6:9], sorted[0])...)
	segmentMap[2] = solver.uniqueChars(segmentMap[5], sorted[0])

	return segmentMap
}

func (solver *Solver) toDigit(segmentMap [7]string, code string) int {
	switch len(code) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	case 5:
		if strings.Contains(code, segmentMap[4]) {
			return 2
		} else if strings.Contains(code, segmentMap[1]) {
			return 5
		} else {
			return 3
		}
	case 6:
		if !strings.Contains(code, segmentMap[3]) {
			return 0
		} else if !strings.Contains(code, segmentMap[2]) {
			return 6
		} else {
			return 9
		}
	}

	panic(fmt.Sprintf("invalid code '%s'", code))
}

func (solver *Solver) sortSegmentsByLength(segments []string) []string {
	var sorted = make([]string, len(segments))
	copy(sorted, segments)
	sort.Slice(sorted, func(i, j int) bool {
		return len(sorted[i]) < len(sorted[j])
	})

	return sorted
}

func (solver *Solver) uniqueChars(str ...string) string {
	var unique, tested string

	for idx, s := range str {
		for i := 0; i < len(s); i++ {
			r := rune(s[i])
			isUnique := true

			if strings.ContainsRune(tested, r) {
				continue
			}

			for i := idx + 1; i < len(str); i++ {
				if strings.ContainsRune(str[i], r) {
					isUnique = false
					break
				}
			}

			if isUnique {
				unique += string(r)
			}
			tested += string(r)
		}
	}

	return unique
}

func (solver *Solver) intersect(str ...string) string {
	var intersection, tested string
	for _, s := range str {

		for i := 0; i < len(s); i++ {
			r := rune(s[i])
			if strings.ContainsRune(tested, r) {
				continue
			}

			tested += string(r)
			intersectRune := true
			for j := 0; j < len(str); j++ {
				if !strings.ContainsRune(str[j], r) {
					intersectRune = false
					break
				}
			}

			if intersectRune {
				intersection += string(r)
			}
		}
	}

	return intersection
}
