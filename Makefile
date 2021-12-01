AOC_PATH=./cmd/aoc/
OUT=./build/

all: run

run:
	go build -o $(OUT) $(AOC_PATH)

win_x64: GOOS=windows
win_x64: GOARCH=64
win_x64: run

win_x32: GOOS=windows
win_x32: GOARCH=386
win_x32: run
