package common

import "strconv"

func ParseIntDecimal(str string) (int, bool) {
	parse, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, false
	}

	return int(parse), true
}

type Stack struct {
	values []interface{}
}

func (s *Stack) Pop() (v interface{}) {
	if len(s.values) == 0 {
		return nil
	}

	v = s.values[0]
	s.values = s.values[1:]
	return v
}

func (s *Stack) Push(v interface{}) {
	s.values = append([]interface{}{v}, s.values...)
}

func (s *Stack) Peak() (v interface{}) {
	return s.values[0]
}

func (s *Stack) Size() int {
	return len(s.values)
}
