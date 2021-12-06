package days

import (
	"fmt"
	"strconv"
	"strings"
)

type fish struct {
	age int
}

type DaySix struct {
	swarm []*fish
}

func (day *DaySix) SolveStarOne(input []string) string {
	return day.parseInput(input).reproduce(80).result()
}

func (day *DaySix) SolveStarTwo(input []string) string {
	return day.parseInput(input).reproduce(160).result()
}

func (day *DaySix) parseInput(input []string) *DaySix {
	if len(input) != 1 {
		panic(fmt.Sprintf("invalid input '%+v'", input))
	}

	fishes := strings.Split(strings.TrimSpace(input[0]), ",")
	day.swarm = make([]*fish, len(fishes))
	for idx, s := range fishes {
		age, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("invalid age '%s'", s))
		}

		day.swarm[idx] = &fish{age: age}
	}

	return day
}

func (day *DaySix) reproduce(days int) *DaySix {
	for i := 0; i < days; i++ {
		currentSwarmLength := len(day.swarm)
		for j := 0; j < currentSwarmLength; j++ {
			if day.swarm[j].age == 0 {
				day.swarm = append(day.swarm, &fish{age: 8})
				day.swarm[j].age = 6
				continue
			}

			day.swarm[j].age -= 1
		}
	}

	return day
}

func (day *DaySix) result() string {
	return fmt.Sprintf("%d", len(day.swarm))
}

func (day *DaySix) String() string {
	result := ""
	for _, f := range day.swarm {
		result += fmt.Sprintf("%d,", f.age)
	}
	return result
}
