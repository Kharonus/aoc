package thirteen

import (
	"fmt"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

type fold struct {
	axis  axis
	value int
}

type axis string

const (
	x axis = "x"
	y axis = "y"
)

type Solver struct {
	dots         []*coordinate
	instructions []*fold
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).fold(1).result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).fold(len(solver.instructions)).result()
}

func (solver *Solver) parseInput(input []string) *Solver {
	solver.dots = make([]*coordinate, 0, len(input))
	emptyLineIndex := 0

	for idx, line := range input {
		if line == "" {
			emptyLineIndex = idx
			break
		}

		solver.dots = append(solver.dots, parseCoordinate(line))
	}

	for _, line := range input[emptyLineIndex+1:] {
		solver.instructions = append(solver.instructions, parseFoldInstruction(line))
	}

	return solver
}

func (solver *Solver) fold(number int) *Solver {
	for i := 0; i < number; i++ {
		f := solver.instructions[0]
		solver.instructions = solver.instructions[1:]

		solver.foldOnce(f)
	}

	return solver
}

func (solver *Solver) foldOnce(f *fold) {
	var dotsOverEdge = make([]*coordinate, 0, len(solver.dots))
	var dotsUnderEdge = make([]*coordinate, 0, len(solver.dots))

	for _, d := range solver.dots {
		switch f.axis {
		case x:
			if d.x < f.value {
				dotsUnderEdge = append(dotsUnderEdge, d)
			} else if d.x > f.value {
				dotsOverEdge = append(dotsOverEdge, d)
			}
		case y:
			if d.y < f.value {
				dotsUnderEdge = append(dotsUnderEdge, d)
			} else if d.y > f.value {
				dotsOverEdge = append(dotsOverEdge, d)
			}
		}
	}

	solver.dots = dotsUnderEdge

	for _, d := range dotsOverEdge {
		switch f.axis {
		case x:
			solver.insertDotIfNotOverlapping(&coordinate{x: d.x - 2*absInt(f.value-d.x), y: d.y})
		case y:
			solver.insertDotIfNotOverlapping(&coordinate{x: d.x, y: d.y - 2*absInt(f.value-d.y)})
		}
	}
}

func (solver *Solver) insertDotIfNotOverlapping(dot *coordinate) {
	for _, d := range solver.dots {
		if d.x == dot.x && d.y == dot.y {
			return
		}
	}

	solver.dots = append(solver.dots, dot)
}

func absInt(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func parseCoordinate(line string) *coordinate {
	values := strings.Split(line, ",")
	if len(values) != 2 {
		panic(fmt.Sprintf("invalid coordinate input '%s'", line))
	}

	x, err := strconv.Atoi(values[0])
	if err != nil {
		panic(fmt.Sprintf("invalid coordinate input '%s'", line))
	}

	y, err := strconv.Atoi(values[1])
	if err != nil {
		panic(fmt.Sprintf("invalid coordinate input '%s'", line))
	}

	return &coordinate{x: x, y: y}
}

func parseFoldInstruction(line string) *fold {
	if strings.Index(line, "fold along ") != 0 {
		panic(fmt.Sprintf("invalid fold instruction '%s'", line))
	}

	f := strings.Split(line[11:], "=")
	value, err := strconv.Atoi(f[1])
	if err != nil {
		panic(fmt.Sprintf("invalid fold instruction '%s'", line))
	}

	switch axis(f[0]) {
	case x:
		return &fold{axis: x, value: value}
	case y:
		return &fold{axis: y, value: value}
	}
	panic(fmt.Sprintf("invalid fold instruction '%s'", line))
}

func (solver *Solver) result() string {
	fmt.Println(solver)
	return fmt.Sprintf("%d", len(solver.dots))
}

func (solver *Solver) String() string {
	var maxX, maxY int
	dotMap := map[coordinate]bool{}

	for _, d := range solver.dots {
		dotMap[*d] = true

		if d.x > maxX {
			maxX = d.x
		}
		if d.y > maxY {
			maxY = d.y
		}
	}

	result := ""
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if dotMap[coordinate{x: j, y: i}] {
				result += "#"
			} else {
				result += "."
			}
		}

		result += "\n"
	}
	return result
}
