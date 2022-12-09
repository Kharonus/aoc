package common

func Sum(a, b int) int {
	return a + b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) (int, int) {
	if a < 0 {
		return a * -1, -1
	}

	return a, 1
}
