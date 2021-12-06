package days

import (
	"fmt"
	"strconv"
	"strings"
)

type DaySix struct {
	swarm [9]int
}

func (day *DaySix) SolveStarOne(input []string) string {
	return day.parseInput(input).reproduce(80).result()
}

func (day *DaySix) SolveStarTwo(input []string) string {
	return day.parseInput(input).reproduce(256).result()
}

func (day *DaySix) parseInput(input []string) *DaySix {
	if len(input) != 1 {
		panic(fmt.Sprintf("invalid input '%+v'", input))
	}

	fishes := strings.Split(strings.TrimSpace(input[0]), ",")
	for _, s := range fishes {
		age, err := strconv.Atoi(s)
		if err != nil || age < 0 || age > 8 {
			panic(fmt.Sprintf("invalid age '%s'", s))
		}

		day.swarm[age] += 1
	}

	return day
}

func (day *DaySix) reproduce(days int) *DaySix {
	for i := 0; i < days; i++ {
		var newSwarm [9]int
		newSwarm[0] = day.swarm[1]
		newSwarm[1] = day.swarm[2]
		newSwarm[2] = day.swarm[3]
		newSwarm[3] = day.swarm[4]
		newSwarm[4] = day.swarm[5]
		newSwarm[5] = day.swarm[6]
		newSwarm[6] = day.swarm[7] + day.swarm[0]
		newSwarm[7] = day.swarm[8]
		newSwarm[8] = day.swarm[0]

		day.swarm = newSwarm
	}

	return day
}

func (day *DaySix) result() string {
	sum := day.swarm[0] + day.swarm[1] + day.swarm[2] + day.swarm[3] +
		day.swarm[4] + day.swarm[5] + day.swarm[6] + day.swarm[7] + day.swarm[8]
	return fmt.Sprintf("%d", sum)
}
