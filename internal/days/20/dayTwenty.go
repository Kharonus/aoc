package twenty

import (
	"fmt"
	"strconv"
)

type step int

const (
	even step = 0
	odd  step = 1
)

type Solver struct {
	enhancementMap map[int]bool
	image          [][]bool
	dx, dy         int
	step           step
}

func (solver *Solver) SolveStarOne(input []string) string {
	return solver.parseInput(input).enhanceUntil(2).result()
}

func (solver *Solver) SolveStarTwo(input []string) string {
	return solver.parseInput(input).enhanceUntil(50).result()
}

func (solver *Solver) parseInput(input []string) *Solver {
	solver.parseEnhancementAlgorithm(input[0])

	solver.step = even
	solver.dy = len(input[2:])
	solver.image = make([][]bool, solver.dy)
	for row, line := range input[2:] {
		solver.dx = len(line)
		solver.image[row] = make([]bool, solver.dx)
		for idx, c := range line {
			if c == '#' {
				solver.image[row][idx] = true
			}
		}
	}

	return solver
}

func (solver *Solver) enhanceUntil(turns int) *Solver {
	for i := 0; i < turns; i++ {
		solver.enhance()
	}

	return solver
}

func (solver *Solver) enhance() *Solver {
	solver.increaseImage()

	enhancedImage := make([][]bool, solver.dy)
	for row := range solver.image {
		enhancedImage[row] = make([]bool, solver.dx)
		for idx := range solver.image[row] {
			enhancedImage[row][idx] = solver.getPixelValue(row, idx)
		}
	}

	if solver.step == even {
		solver.step = odd
	} else {
		solver.step = even
	}

	solver.image = enhancedImage
	return solver
}

func (solver *Solver) getPixelValue(i, j int) bool {
	grid := [3][3]bool{}
	if solver.step == odd && solver.enhancementMap[0] {
		// TODO: this very behavior is still not correct, try to fix it
		// after every odd step the infinite image is lit instead of dark if the first pixel
		// of the enhancement map is light
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				grid[i][j] = true
			}
		}
	}

	grid[1][1] = solver.image[i][j]

	if i > 0 && j > 0 {
		grid[0][0] = solver.image[i-1][j-1]
	}
	if i > 0 {
		grid[0][1] = solver.image[i-1][j]
	}
	if j > 0 {
		grid[1][0] = solver.image[i][j-1]
	}
	if i > 0 && j < solver.dx-1 {
		grid[0][2] = solver.image[i-1][j+1]
	}
	if i < solver.dy-1 {
		grid[2][1] = solver.image[i+1][j]
	}
	if j < solver.dx-1 {
		grid[1][2] = solver.image[i][j+1]
	}
	if i < solver.dy-1 && j > 0 {
		grid[2][0] = solver.image[i+1][j-1]
	}
	if i < solver.dy-1 && j < solver.dx-1 {
		grid[2][2] = solver.image[i+1][j+1]
	}

	binary := ""

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch grid[i][j] {
			case true:
				binary += "1"
			case false:
				binary += "0"
			}
		}
	}

	value, _ := strconv.ParseInt(binary, 2, 16)
	return solver.enhancementMap[int(value)]
}

func (solver *Solver) increaseImage() {
	newDx := solver.dx + 2
	newDy := solver.dy + 2
	image := make([][]bool, newDy)
	for i := 0; i < newDy; i++ {
		image[i] = make([]bool, newDx)
		for j := 0; j < len(image[i]); j++ {
			if i > 0 && j > 0 && i < solver.dy+1 && j < solver.dx+1 {
				image[i][j] = solver.image[i-1][j-1]
			}
		}
	}

	solver.dx = newDx
	solver.dy = newDy
	solver.image = image
}

func (solver *Solver) parseEnhancementAlgorithm(line string) {
	solver.enhancementMap = map[int]bool{}
	for idx, c := range line {
		switch c {
		case '.':
			solver.enhancementMap[idx] = false
		case '#':
			solver.enhancementMap[idx] = true
		}
	}
}

func (solver *Solver) printImage() {
	result := ""
	for _, row := range solver.image {
		for _, pixel := range row {
			switch pixel {
			case true:
				result += "#"
			case false:
				result += "."
			}
		}
		result += "\n"
	}

	fmt.Println(result)
}

func (solver *Solver) result() string {
	count := 0
	for row := range solver.image {
		for _, p := range solver.image[row] {
			if p {
				count++
			}
		}
	}

	return fmt.Sprintf("%d", count)
}
