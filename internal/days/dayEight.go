package days

import (
	"fmt"
	"sort"
	"strings"
)

type DayEight struct {
	digitsUnique []string
}

func (day *DayEight) SolveStarOne(input []string) string {
	return day.parseInput(input).resultUniqueDigits()
}

func (day *DayEight) SolveStarTwo(input []string) string {
	return day.parseInput(input).resultUniqueDigits()
}

func (day *DayEight) parseInput(input []string) *DayEight {
	day.digitsUnique = make([]string, 0, len(input)*4)
	for _, s := range input {
		segments, digits := day.parseLine(s)
		day.deduceSegments(segments)

		for _, d := range digits {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				day.digitsUnique = append(day.digitsUnique, d)
			}
		}
	}

	return day
}

func (day DayEight) parseLine(line string) (segments, digits []string) {
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

func (day *DayEight) deduceSegments(segments []string) {
	var segmentMap [7]string
	sorted := day.sortSegmentsByLength(segments)

	sorted[0]
}

func (day DayEight) sortSegmentsByLength(segments []string) []string {
	var sorted = make([]string, len(segments))
	copy(sorted, segments)
	sort.Slice(sorted, func(i, j int) bool {
		return len(sorted[i]) < len(sorted[j])
	})

	return sorted
}

func (day *DayEight) uniqueChars(str... string) string {
	var unique, duplicate string

	for _, s := range str {
		for i := 0; i < len(s); i++ {
			if !strings.Contains(unique, s[i:i]) {
				unique += s[i:i]
			} else {

			}
		}
	}

	for i := 0; i < len(str); i++ {

	}
}
