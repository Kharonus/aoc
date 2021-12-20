package nineteen

import (
	"fmt"
	"strconv"
	"strings"
)

type location struct {
	x, y, z int
}

type overlap struct {
	transformation [4][4]int
	from           *scanner
	to             *scanner
}

type scanner struct {
	beacons []*location
}

type Solver struct {
	scanners []*scanner
	beacons  []*location
}

func (solver *Solver) SolveStarOne(input []string) string {
	solver.parseInput(input)
	return "What, are you impatient? We do not even approached this far in December 2021 ..."
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return "What, are you impatient? We do not even approached this far in December 2021 ..."
}

func (solver *Solver) parseInput(input []string) *Solver {
	var beacons []*location
	for _, line := range input {
		if strings.Index(line, "---") == 0 {
			beacons = []*location{}
			continue
		}

		if line == "" {
			solver.scanners = append(solver.scanners, &scanner{beacons: beacons})
			continue
		}

		beacons = append(beacons, parseLocation(line))
	}

	return solver
}

func (solver *Solver) findOverlap(from, to *scanner) *overlap {



}

func computeTransformation(v1, v2 *location) [4][4]int {

}

func parseLocation(line string) *location {
	split := strings.Split(line, ",")
	x, err := strconv.Atoi(split[0])
	if err != nil {
		panic(fmt.Sprintf("invalid coordinate line '%s'", line))
	}
	y, err := strconv.Atoi(split[1])
	if err != nil {
		panic(fmt.Sprintf("invalid coordinate line '%s'", line))
	}
	z, err := strconv.Atoi(split[2])
	if err != nil {
		panic(fmt.Sprintf("invalid coordinate line '%s'", line))
	}

	return &location{x: x, y: y, z: z}
}
