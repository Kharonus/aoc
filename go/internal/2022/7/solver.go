package _7

import (
	"fmt"
	"github.com/Kharonus/aoc/internal/common"
	"regexp"
	"strconv"
)

type file struct {
	name string
	size int
}

type directory struct {
	name      string
	parent    *directory
	subs      []*directory
	files     []*file
	totalSize int
}

type Solver struct {
	root, current *directory
}

func (day *Solver) SolveStarOne(input []string) string {
	size := day.parseInput(input).calculateTotalSizes().root.sumDirectoriesWithLessThen(100000)
	return strconv.Itoa(size)
}

func (day *Solver) SolveStarTwo(input []string) string {
	size := day.parseInput(input).
		calculateTotalSizes().
		findDirectoryToDelete(70000000, 30000000).
		totalSize
	return strconv.Itoa(size)
}

func (day *Solver) parseInput(input []string) *Solver {
	day.root = &directory{name: "/"}

	for idx, line := range input {
		if !isCommand(line) {
			continue
		}

		switch line[2:4] {
		case "cd":
			day.executeChangeDirectory(line[5:])
		case "ls":
			day.executeList(input[idx+1:])
		default:
			panic(fmt.Sprintf("Unknown command '%s'.", line))
		}
	}

	return day
}

func (day *Solver) executeList(input []string) {
	for _, line := range input {
		if isCommand(line) {
			return
		}

		if ok, dir := isDirectory(line); ok {
			dir.parent = day.current
			day.current.subs = append(day.current.subs, dir)
			continue
		}

		if ok, f := isFile(line); ok {
			day.current.files = append(day.current.files, f)
		}
	}
}

func (day *Solver) executeChangeDirectory(target string) {
	switch target {
	case "/":
		day.current = day.root
	case "..":
		day.current = day.current.parent
	default:
		var t *directory
		for _, sub := range day.current.subs {
			if sub.name == target {
				t = sub
				break
			}
		}

		if t == nil {
			panic(fmt.Sprintf("Invalid cd target: '%s', no such directory found.", target))
		}

		day.current = t
	}
}

func (day *Solver) calculateTotalSizes() *Solver {
	day.root.calculateTotalSize()
	return day
}

func (dir *directory) calculateTotalSize() {
	for _, sub := range dir.subs {
		sub.calculateTotalSize()
	}

	subsSize := common.Reduce(dir.subs, func(sum int, sub *directory) int {
		return sum + sub.totalSize
	}, 0)

	filesSize := common.Reduce(dir.files, func(sum int, f *file) int {
		return sum + f.size
	}, 0)

	dir.totalSize = subsSize + filesSize
}

func (day *Solver) findDirectoryToDelete(total, needed int) *directory {
	free := total - day.root.totalSize
	toDelete := needed - free
	if toDelete < 0 {
		panic(fmt.Sprintf("You have enough space ('%d')", free))
	}

	toCheck := []*directory{day.root}
	var deletionTarget *directory

	for len(toCheck) > 0 {
		dir := toCheck[0]
		toCheck = toCheck[1:]

		if dir.totalSize > toDelete {
			if deletionTarget == nil || deletionTarget.totalSize > dir.totalSize {
				deletionTarget = dir
			}
			toCheck = append(toCheck, dir.subs...)
		}
	}

	return deletionTarget
}

func (dir *directory) sumDirectoriesWithLessThen(size int) int {
	var sum int
	for _, sub := range dir.subs {
		sum += sub.sumDirectoriesWithLessThen(size)
	}

	if dir.totalSize < size {
		sum += dir.totalSize
	}

	return sum
}

func isCommand(line string) bool {
	return line[0] == '$'
}

func isDirectory(line string) (bool, *directory) {
	if line[:3] == "dir" {
		return true, &directory{name: line[4:]}
	}

	return false, nil
}

func isFile(line string) (bool, *file) {

	r := regexp.MustCompile(`[0-9]+`)
	match := r.FindString(line)
	index := r.FindStringIndex(line)

	if match != "" {
		size, _ := strconv.ParseInt(match, 10, 32)

		return true, &file{
			name: line[index[1]:],
			size: int(size),
		}
	}

	return false, nil
}
