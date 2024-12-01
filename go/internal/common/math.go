package common

func Sum(a, b int) int {
	return a + b
}

func Product(a, b int) int {
	return a * b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
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

func Diff(a, b int) int {
	abs, _ := Abs(a - b)
	return abs
}

func Manhattan(a, b [2]int) int {
	x, _ := Abs(a[0] - b[0])
	y, _ := Abs(a[1] - b[1])
	return x + y
}
