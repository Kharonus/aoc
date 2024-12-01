package internal

import (
	"bufio"
	"os"
)

func ReadFileLineByLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	return result
}
