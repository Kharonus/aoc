package common

import "strconv"

func ParseIntDecimal(str string) (int, bool) {
	parse, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, false
	}

	return int(parse), true
}
