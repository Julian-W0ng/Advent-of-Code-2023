package main

import (
	"os"
	"strconv"
	"strings"
)

const TIME int = 0
const DISTANCE int = 1

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// remove all spaces
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			input = append(input[:i], input[i+1:]...)
			i--
		}
	}

	lines := strings.Split(string(input), "\n")

	time, _ := strconv.Atoi(strings.Split(lines[0], ":")[1])
	distance, _ := strconv.Atoi(strings.Split(lines[1], ":")[1])

	waysToWin := 0

	for i := 0; i <= time; i++ {
		distanceCovered := i * (time - i)
		if distanceCovered > distance {
			waysToWin++
		}
	}

	println(waysToWin)
}
