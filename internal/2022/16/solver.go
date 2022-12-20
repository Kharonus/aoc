package _16

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Kharonus/aoc/internal/common"
)

type valve struct {
	name      string
	flow      int
	neighbors map[*valve]int
}

type valveOld struct {
	name     string
	flowRate int
	tunnels  []string
}

type Solver struct {
	valvesOld       map[string]*valveOld
	options         []*option
	openValves      map[string]struct{}
	currentPosition string

	valves        map[string]*valve
	startPosition string
}

type option struct {
	openValves       map[string]struct{}
	releasedPressure int
	remainingMinutes int
	currentPosition  string
	order            []string
}

func (day *Solver) SolveStarOne(input []string) string {
	defer func() {
		if r := recover(); r != nil {
			panic("und nu?")
		}
	}()

	released := day.parseInput(input).preprocessValveGraph().releaseMaxPressure(30)
	return strconv.Itoa(released)
}

func (day *Solver) SolveStarTwo(input []string) string {
	return "What, are you impatient? We didn't even reach this date yet."
}

func (day *Solver) parseInput(input []string) *Solver {
	day.valvesOld = map[string]*valveOld{}
	day.valves = map[string]*valve{}
	day.openValves = map[string]struct{}{}
	day.currentPosition = "AA"
	day.startPosition = "AA"

	for _, line := range input {
		day.parseLine(line)

		name, val := parseLineOld(line)
		day.valvesOld[name] = val
	}

	return day
}

func parseLineOld(line string) (string, *valveOld) {
	valves := regexp.MustCompile(`[A-Z]{2}`).FindAllString(line, -1)
	rate, _ := common.ParseIntDecimal(regexp.MustCompile(`[0-9]+`).FindString(line))

	return valves[0], &valveOld{name: valves[0], flowRate: rate, tunnels: valves[1:]}
}

func (day *Solver) parseLine(line string) *valve {
	valveNames := regexp.MustCompile(`[A-Z]{2}`).FindAllString(line, -1)
	rate, _ := common.ParseIntDecimal(regexp.MustCompile(`[0-9]+`).FindString(line))

	var result *valve
	for idx, name := range valveNames {
		v := day.getValve(name)

		switch idx {
		case 0:
			v.flow = rate
			v.neighbors = map[*valve]int{}
			result = v
		default:
			result.neighbors[v] = 1
		}
	}

	return result
}

func (day *Solver) getValve(name string) *valve {
	if v, ok := day.valves[name]; ok {
		return v
	}

	v := &valve{name: name}
	day.valves[name] = v
	return v
}

func (day *Solver) releaseMaxPressureOld(minutes int) int {
	var pastMinutes, releasedPressure int

	for pastMinutes < minutes {
		pastMinutes += day.moveToMostValuableValve(minutes - pastMinutes)
		if day.currentPosition == "" {
			break
		}

		pastMinutes += 1
		current := day.valvesOld[day.currentPosition]
		day.openValves[current.name] = struct{}{}
		releasedPressure += current.flowRate * (minutes - pastMinutes)
	}

	return releasedPressure
}

func (day *Solver) releaseMaxPressure(minutes int) int {
	var options = []*option{
		{
			openValves:       map[string]struct{}{},
			releasedPressure: 0,
			remainingMinutes: minutes,
			currentPosition:  day.startPosition,
		},
	}

	var bestOption = &option{releasedPressure: 0}

	for len(options) > 0 {
		currentOpt := options[0]
		options = options[1:]

		newOpts := day.expandOption(currentOpt)
		if len(newOpts) == 0 && currentOpt.releasedPressure > bestOption.releasedPressure {
			bestOption = currentOpt
		}
		options = append(options, newOpts...)
	}

	return bestOption.releasedPressure
}

func (day *Solver) expandOption(opt *option) []*option {
	var moreOptions []*option
	v := day.valves[opt.currentPosition]

	// Do not search for neighbors, instead search for all open values
	for val, dist := range v.neighbors {
		if _, ok := opt.openValves[val.name]; ok {
			continue
		}

		var openValves = map[string]struct{}{}
		openValves[val.name] = struct{}{}
		for name := range opt.openValves {
			openValves[name] = struct{}{}
		}

		newOption := &option{
			openValves:       openValves,
			releasedPressure: opt.releasedPressure + val.flow*(opt.remainingMinutes-dist),
			remainingMinutes: opt.remainingMinutes - (dist + 1),
			currentPosition:  val.name,
			order:            append(opt.order, val.name),
		}

		moreOptions = append(moreOptions, newOption)
	}

	return moreOptions
}

func (day *Solver) preprocessValveGraph() *Solver {
	var collapsedValves []*valve

	for _, v := range day.valves {
		if v.flow == 0 && v.name != day.startPosition {
			day.collapseValveNode(v)
			collapsedValves = append(collapsedValves, v)
		}
	}

	for _, v := range collapsedValves {
		delete(day.valves, v.name)
	}

	return day
}

func (day *Solver) collapseValveNode(val *valve) {
	for v, d1 := range val.neighbors {
		for target, d2 := range val.neighbors {
			if v == target {
				continue
			}

			delete(v.neighbors, val)
			if dist, ok := v.neighbors[target]; !ok || dist > d1+d2 {
				v.neighbors[target] = d1 + d2
			}
		}
	}
}

func (day *Solver) eliminateOptions() {
	var validOptions []*option

	for i, opt1 := range day.options {
		isValid := false

		for j, opt2 := range day.options {
			if i == j {
				continue
			}

			if compare(opt1, opt2) {
				isValid = true
				break
			}
		}

		if isValid {
			validOptions = append(validOptions, opt1)
		}
	}

	day.options = validOptions
}

func compare(opt1, opt2 *option) bool {
	return opt2.remainingMinutes <= opt1.remainingMinutes || opt2.releasedPressure >= opt2.releasedPressure
}

func (day *Solver) moveToMostValuableValve(remainingMinutes int) (pastMinutes int) {
	var mostValuableValve string
	var possiblePressureRelease, minutes int

	calcPressure := func(v *valveOld, path int) int {
		return v.flowRate * (remainingMinutes - (path + 1))
	}

	for name, val := range day.valvesOld {
		if _, ok := day.openValves[name]; ok {
			continue
		}

		path := day.shortestPathTo(name)
		pressure := calcPressure(val, path)

		if pressure > possiblePressureRelease {
			possiblePressureRelease = pressure
			minutes = path
			mostValuableValve = name
		}
	}

	day.currentPosition = mostValuableValve
	return minutes
}

func (day *Solver) expandOptions() {
	//for _, opt := range day.options {
	//	current := opt.currentPosition
	//}
}

func (day *Solver) expandOptionOld(opt *option) []*option {
	var validOptions []*option
	current := day.valvesOld[opt.currentPosition]

	type tuple struct {
		name string
		dist int
	}

	var next []tuple

	for _, t := range current.tunnels {
		next = append(next, tuple{name: t, dist: 1})
	}

	//for len(next) > 0 {
	//	n := day.valvesOld[next[0]]
	//	next = next[1:]
	//
	//}

	return validOptions
}

func (day *Solver) shortestPathTo(target string) int {
	var next = []string{day.currentPosition}
	var hopMap = map[string]int{day.currentPosition: 0}

	for len(next) > 0 {
		current := day.valvesOld[next[0]]
		next = next[1:]

		for _, t := range current.tunnels {
			if t == target {
				return hopMap[current.name] + 1
			}

			if _, ok := hopMap[t]; ok {
				continue
			}

			next = append(next, t)
			hopMap[t] = hopMap[current.name] + 1
		}
	}

	panic(fmt.Sprintf("No path found from '%s' to '%s'.", day.currentPosition, target))
}
