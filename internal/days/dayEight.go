package days

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

type DayEight struct {
	digitsUnique []string
	lines        []*codedLine
}

func (day *DayEight) SolveStarOne(input []string) string {
	return day.parseInput(input).resultUniqueDigits()
}

func (day *DayEight) SolveStarTwo(input []string) string {
	return day.parseInput(input).resultDecodedSum()
}

func (day *DayEight) parseInput(input []string) *DayEight {
	day.digitsUnique = make([]string, 0, len(input)*4)
	day.lines = make([]*codedLine, 0, len(input))
	for _, s := range input {
		segments, digits := day.parseLine(s)

		day.lines = append(day.lines, &codedLine{
			digits:     digits,
			segmentMap: day.deduceSegments(segments),
		})

		for _, d := range digits {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				day.digitsUnique = append(day.digitsUnique, d)
			}
		}
	}

	return day
}

func (day *DayEight) parseLine(line string) (segments, digits []string) {
	split := strings.Split(strings.TrimSpace(line), "|")
	if len(split) != 2 {
		panic(fmt.Sprintf("invalid line '%s'", line))
	}

	segments = strings.Split(strings.TrimSpace(split[0]), " ")
	digits = strings.Split(strings.TrimSpace(split[1]), " ")
	return
}

func (day *DayEight) resultUniqueDigits() string {
	return fmt.Sprintf("%d", len(day.digitsUnique))
}

func (day *DayEight) resultDecodedSum() string {
	sum := 0
	for _, line := range day.lines {
		number := 0
		for i := 0; i < len(line.digits); i++ {
			number += int(math.Pow10(3-i)) * day.toDigit(line.segmentMap, line.digits[i])
		}
		sum += number
	}

	return fmt.Sprintf("%d", sum)
}

func (day *DayEight) deduceSegments(segments []string) [7]string {
	var segmentMap [7]string
	sorted := day.sortSegmentsByLength(segments)

	leftSide := day.uniqueChars(sorted[3:6]...)
	topLeftMiddle := day.uniqueChars(sorted[0], sorted[2])

	segmentMap[0] = day.uniqueChars(sorted[0], sorted[1])
	segmentMap[1] = day.intersect(leftSide, topLeftMiddle)
	segmentMap[3] = day.uniqueChars(segmentMap[1], topLeftMiddle)
	segmentMap[4] = day.uniqueChars(leftSide, segmentMap[1])
	segmentMap[6] = day.uniqueChars(leftSide, segmentMap[3], sorted[1], sorted[9])

	segmentMap[5] = day.intersect(append(sorted[6:9], sorted[0])...)
	segmentMap[2] = day.uniqueChars(segmentMap[5], sorted[0])

	return segmentMap
}

func (day *DayEight) toDigit(segmentMap [7]string, code string) int {
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

func (day *DayEight) sortSegmentsByLength(segments []string) []string {
	var sorted = make([]string, len(segments))
	copy(sorted, segments)
	sort.Slice(sorted, func(i, j int) bool {
		return len(sorted[i]) < len(sorted[j])
	})

	return sorted
}

func (day *DayEight) uniqueChars(str ...string) string {
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

func (day *DayEight) intersect(str ...string) string {
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
