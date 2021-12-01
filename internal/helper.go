package internal

import "strconv"

func StringSliceToIntSlice(input []string) ([]int, error) {
	var result = make([]int, len(input))
	for idx, s := range input {
		value, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		result[idx] = value
	}

	return result, nil
}
