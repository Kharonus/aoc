package days

import (
	"fmt"
	"strconv"
	"strings"
)

type DayTwo struct {
	reach int
	depth int
	aim   int
}

var orderFunctions = map[string]func(day *DayTwo, value int){
	"forward": func(day *DayTwo, value int) { day.reach += value },
	"up":      func(day *DayTwo, value int) { day.depth -= value },
	"down":    func(day *DayTwo, value int) { day.depth += value },
}

var orderFunctionsWithAim = map[string]func(day *DayTwo, value int){
	"forward": func(day *DayTwo, value int) {
		day.reach += value
		day.depth += day.aim * value
	},
	"up":   func(day *DayTwo, value int) { day.aim -= value },
	"down": func(day *DayTwo, value int) { day.aim += value },
}

func (day *DayTwo) SolveStarOne(input []string) string {
	for _, s := range input {
		day.interpretInput(s, false)
	}

	return day.result()
}

func (day *DayTwo) SolveStarTwo(input []string) string {
	for _, s := range input {
		day.interpretInput(s, true)
	}

	return day.result()
}

func (day *DayTwo) interpretInput(order string, withAim bool) *DayTwo {
	splits := strings.Split(strings.TrimSpace(order), " ")
	if len(splits) != 2 {
		panic(fmt.Sprintf("invalid order '%s'", order))
	}

	value, err := strconv.Atoi(splits[1])
	if err != nil {
		panic(fmt.Sprintf("invalid order %s", order))
	}

	var orderFn func(day *DayTwo, value int)
	if withAim {
		orderFn = orderFunctionsWithAim[splits[0]]
	} else {
		orderFn = orderFunctions[splits[0]]
	}

	if orderFn == nil {
		panic(fmt.Sprintf("invalid order %s", order))
	}

	orderFn(day, value)
	return day
}

func (day *DayTwo) result() string {
	return fmt.Sprintf("%d", day.reach*day.depth)
}
