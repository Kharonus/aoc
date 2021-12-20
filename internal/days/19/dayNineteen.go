package nineteen

import (
	"fmt"
	"strconv"
	"strings"
)

type transformation struct {
	rotation    [3][3]int
	translation *vector
}

type vector struct {
	x, y, z int
}

type overlap struct {
	transformation *transformation
	from           *scanner
	to             *scanner
}

type scanner struct {
	beacons        []*vector
	transformation *transformation
}

type Solver struct {
	scanners []*scanner
	beacons  []*vector
}

var possibleRotations = [][3][3]int{
	{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	},
	{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	},
	{
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	},
	{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, -1},
	},
	{
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	},
	{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, -1},
	},
	{
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	},
	{
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	},
	{
		{1, 0, 0},
		{0, 0, 1},
		{0, -1, 0},
	},
	{
		{-1, 0, 0},
		{0, 0, 1},
		{0, -1, 0},
	},
	{
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	},
	{
		{-1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	},
	{
		{0, 0, -1},
		{0, 1, 0},
		{1, 0, 0},
	},
	{
		{0, 0, -1},
		{0, -1, 0},
		{1, 0, 0},
	},
	{
		{0, 0, 1},
		{0, 1, 0},
		{-1, 0, 0},
	},
	{
		{0, 0, 1},
		{0, -1, 0},
		{-1, 0, 0},
	},
	{
		{0, 1, 0},
		{-1, 0, 0},
		{0, 0, 1},
	},
	{
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, 1},
	},
	{
		{0, 1, 0},
		{-1, 0, 0},
		{0, 0, -1},
	},
	{
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, -1},
	},
}

func (solver *Solver) SolveStarOne(input []string) string {
	solver.parseInput(input).convertBeaconsToBase()
	return "What, are you impatient? We do not even approached this far in December 2021 ..."
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return "What, are you impatient? We do not even approached this far in December 2021 ..."
}

func (solver *Solver) convertBeaconsToBase() *Solver {
	var overlaps []*overlap
	for idx, s := range solver.scanners {
		for i := idx + 1; i < len(solver.scanners); i++ {
			intersection := solver.findOverlap(s, solver.scanners[i])
			if intersection != nil {
				overlaps = append(overlaps, intersection)
			}
		}
	}

	// scanner 0 is the base system
	solver.scanners[0].transformation = &transformation{
		rotation: [3][3]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		translation: &vector{0, 0, 0},
	}

	scannerWithTransformation := make([]*scanner, 0, len(solver.scanners))
	scannerWithTransformation = append(scannerWithTransformation, solver.scanners[0])

	for len(overlaps) > 0 {
		var o *overlap
		for idx, laps := range overlaps {
			// TODO: from and to can be exchanged !!!
			if containsScanner(scannerWithTransformation, laps.from) &&
				!containsScanner(scannerWithTransformation, laps.to) {
				o = laps

				if idx == len(overlaps)-1 {
					overlaps = overlaps[:idx]
				} else {
					overlaps = append(overlaps[:idx], overlaps[idx+1:]...)
				}
				break
			}
		}

		o.to.transformation = &transformation{
			rotation:    multiply(o.from.transformation.rotation, o.transformation.rotation),
			translation: o.from.transformation.translation.add(o.transformation.translation),
		}
		scannerWithTransformation = append(scannerWithTransformation, o.to)
	}

	return solver
}

func containsScanner(list []*scanner, scan *scanner) bool {
	for _, s := range list {
		if s == scan {
			return true
		}
	}
	return false
}

func (solver *Solver) parseInput(input []string) *Solver {
	var beacons []*vector
	for _, line := range input {
		if strings.Index(line, "---") == 0 {
			beacons = []*vector{}
			continue
		}

		if line == "" {
			solver.scanners = append(solver.scanners, &scanner{beacons: beacons})
			continue
		}

		beacons = append(beacons, parseLocation(line))
	}
	solver.scanners = append(solver.scanners, &scanner{beacons: beacons})

	return solver
}

func (solver *Solver) findOverlap(from, to *scanner) *overlap {
	translationMap := map[vector]int{}
	for _, rotation := range generateRotations() {
		for _, fromBeacon := range from.beacons {
			for _, toBeacon := range to.beacons {
				translation := fromBeacon.subtract(toBeacon.rotate(transpose(rotation)))
				translationMap[*translation] += 1

				if translationMap[*translation] >= 12 {
					return &overlap{
						transformation: &transformation{
							rotation:    rotation,
							translation: translation,
						},
						from: from,
						to:   to,
					}
				}
			}
		}
	}

	return nil
	//return &overlap{
	//	transformation: nil,
	//	from:           from,
	//	to:             to,
	//}
}

func generateRotations() [][3][3]int {
	a := [3]int{1, 0, 0}
	b := [3]int{0, 1, 0}
	c := [3]int{0, 0, 1}

	negate := func(v [3]int) [3]int {
		return [3]int{v[0] * -1, v[1] * -1, v[2] * -1}
	}

	return [][3][3]int{
		{a, b, c},
		{a, c, b},
		{b, a, c},
		{b, c, a},
		{c, b, a},
		{c, a, b},
		{negate(a), b, c},
		{negate(a), c, b},
		{b, negate(a), c},
		{b, c, negate(a)},
		{c, b, negate(a)},
		{c, negate(a), b},
		{a, negate(b), c},
		{a, c, negate(b)},
		{negate(b), a, c},
		{negate(b), c, a},
		{c, negate(b), a},
		{c, a, negate(b)},
		{a, b, negate(c)},
		{a, negate(c), b},
		{b, a, negate(c)},
		{b, negate(c), a},
		{negate(c), b, a},
		{negate(c), a, b},
		{negate(a), negate(b), c},
		{negate(a), c, negate(b)},
		{negate(b), negate(a), c},
		{negate(b), c, negate(a)},
		{c, negate(b), negate(a)},
		{c, negate(a), negate(b)},
		{negate(a), b, negate(c)},
		{negate(a), negate(c), b},
		{b, negate(a), negate(c)},
		{b, negate(c), negate(a)},
		{negate(c), b, negate(a)},
		{negate(c), negate(a), b},
		{a, negate(b), negate(c)},
		{a, negate(c), negate(b)},
		{negate(b), a, negate(c)},
		{negate(b), negate(c), a},
		{negate(c), negate(b), a},
		{negate(c), a, negate(b)},
		{negate(a), negate(b), negate(c)},
		{negate(a), negate(c), negate(b)},
		{negate(b), negate(a), negate(c)},
		{negate(b), negate(c), negate(a)},
		{negate(c), negate(b), negate(a)},
		{negate(c), negate(a), negate(b)},
	}
}

func computeTransformation(v1, t1, v2, t2 *vector) *transformation {
	for _, r := range possibleRotations {
		translation1 := v1.subtract(v2.rotate(r))
		translation2 := t1.subtract(t2.rotate(r))

		if translation1.equals(translation2) {
			return &transformation{
				rotation:    r,
				translation: translation1,
			}
		}
	}

	return nil
}

func printR(r [3][3]int) {
	result := ""
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if r[i][j] == -1 {
				result += "-1,"
			} else {
				result += fmt.Sprintf(" %d,", r[i][j])
			}
		}
		result += "\n"
	}

	fmt.Println(result)
}

func (v *vector) subtract(vec *vector) *vector {
	return &vector{
		x: v.x - vec.x,
		y: v.y - vec.y,
		z: v.z - vec.z,
	}
}

func (v *vector) add(vec *vector) *vector {
	return &vector{
		x: v.x + vec.x,
		y: v.y + vec.y,
		z: v.z + vec.z,
	}
}

func (v *vector) equals(vec *vector) bool {
	return v.x == vec.x && v.y == vec.y && v.z == vec.z
}

func (v *vector) scalar(s float64) *vector {
	return &vector{
		x: int(float64(v.x) * s),
		y: int(float64(v.y) * s),
		z: int(float64(v.z) * s),
	}
}

func (v *vector) transform(t *transformation) *vector {
	translated := v.subtract(t.translation)

	return translated.rotate(t.rotation)
}

func (v *vector) rotate(rotation [3][3]int) *vector {
	return &vector{
		x: v.x*rotation[0][0] + v.y*rotation[0][1] + v.z*rotation[0][2],
		y: v.x*rotation[1][0] + v.y*rotation[1][1] + v.z*rotation[1][2],
		z: v.x*rotation[2][0] + v.y*rotation[2][1] + v.z*rotation[2][2],
	}
}

func transpose(rotation [3][3]int) [3][3]int {
	return [3][3]int{
		{rotation[0][0], rotation[1][0], rotation[2][0]},
		{rotation[0][1], rotation[1][1], rotation[2][1]},
		{rotation[0][2], rotation[1][2], rotation[2][2]},
	}
}

func multiply(m1, m2 [3][3]int) [3][3]int {
	return [3][3]int{
		{m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0] + m1[0][2]*m2[2][0], m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1] + m1[0][2]*m2[2][1], m1[0][0]*m2[0][2] + m1[0][1]*m2[1][2] + m1[0][2]*m2[2][2]},
		{m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0] + m1[1][2]*m2[2][0], m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1] + m1[1][2]*m2[2][1], m1[1][0]*m2[0][2] + m1[1][1]*m2[1][2] + m1[1][2]*m2[2][2]},
		{m1[2][0]*m2[0][0] + m1[2][1]*m2[1][0] + m1[2][2]*m2[2][0], m1[2][0]*m2[0][1] + m1[2][1]*m2[1][1] + m1[2][2]*m2[2][1], m1[2][0]*m2[0][2] + m1[2][1]*m2[1][2] + m1[2][2]*m2[2][2]},
	}
}

func parseLocation(line string) *vector {
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

	return &vector{x: x, y: y, z: z}
}
