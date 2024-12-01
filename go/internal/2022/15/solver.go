package _15

import (
	"math"
	"regexp"
	"strconv"

	"github.com/Kharonus/aoc/internal/common"
)

type location int

type coord [2]int

const (
	beacon location = iota
	sensor
	uncharted
)

type sensorBeacon struct {
	sensor, beacon coord
}

type Solver struct {
	tunnels map[coord]location
	sensors []*sensorBeacon
}

func (day *Solver) SolveStarOne(input []string) string {
	coveredLocations := day.parseInput(input).getCoveredLocationsOfLine(2000000)
	return strconv.Itoa(coveredLocations)
}

func (day *Solver) SolveStarTwo(input []string) string {
	tuningFrequency := day.parseInput(input).getDistressBeacon(0, 4000000)
	return strconv.Itoa(tuningFrequency)
}

func (day *Solver) parseInput(input []string) *Solver {
	day.tunnels = map[coord]location{}

	for _, line := range input {
		s, b := parseLine(line)
		day.sensors = append(day.sensors, &sensorBeacon{sensor: s, beacon: b})
		day.tunnels[s] = sensor
		day.tunnels[b] = beacon
	}

	return day
}

func parseLine(line string) (sensor, beacon coord) {
	numbers := regexp.MustCompile(`-?[0-9]+`).FindAllString(line, 4)
	sx, _ := common.ParseIntDecimal(numbers[0])
	sy, _ := common.ParseIntDecimal(numbers[1])
	bx, _ := common.ParseIntDecimal(numbers[2])
	by, _ := common.ParseIntDecimal(numbers[3])

	return coord{sx, sy}, coord{bx, by}
}

func (day *Solver) getLocation(c coord) location {
	if val, ok := day.tunnels[c]; ok {
		return val
	}
	return uncharted
}

func (day *Solver) getCoveredLocationsOfLine(y int) int {
	var covered int

	minX, maxX, _, _ := day.getMinMax()
	for x := minX; x <= maxX; x++ {
		c := coord{x, y}
		loc := day.getLocation(c)
		if loc == beacon || loc == sensor {
			continue
		}

		for _, sb := range day.sensors {
			if isCovered(sb, c) {
				covered += 1
				break
			}
		}
	}

	return covered
}

func isCovered(s *sensorBeacon, c coord) bool {
	dist1 := common.Manhattan(s.sensor, s.beacon)
	dist2 := common.Manhattan(s.sensor, c)

	return dist2 <= dist1
}

func (day *Solver) getDistressBeacon(min, max int) int {
	calcTuningFrequency := func(c coord) int {
		return c[0]*4000000 + c[1]
	}

	for x := min; x <= max; x++ {
		for y := min; y <= max; y++ {
			c := coord{x, y}
			var covered = false

			for _, s := range day.sensors {
				covered = covered || isCovered(s, c)
			}

			if !covered {
				return calcTuningFrequency(c)
			}
		}
	}

	panic("No distress beacon found.")
}

func (day *Solver) getMinMax() (int, int, int, int) {
	var minx, miny = math.MaxInt, math.MaxInt
	var maxx, maxy = math.MinInt, math.MinInt

	for _, pair := range day.sensors {
		x1, x2, y1, y2 := getMinMax(pair.sensor, pair.beacon)
		minx = common.Min(minx, x1)
		maxx = common.Max(maxx, x2)
		miny = common.Min(miny, y1)
		maxy = common.Max(maxy, y2)
	}

	return minx, maxx, miny, maxy
}

func getMinMax(s, b coord) (int, int, int, int) {
	dist := common.Manhattan(s, b)
	return s[0] - dist, s[0] + dist, s[1] - dist, s[1] + dist
}
