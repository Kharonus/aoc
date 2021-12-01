package internal

type IDaySolver interface {
	SolveStarOne(input []string) (string, error)
	SolveStarTwo(input []string) (string, error)
}
