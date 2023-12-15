package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func computeNextDiff(diff []int) []int {
	var nextDiff []int = make([]int, len(diff)-1)
	for i := 0; i < len(diff)-1; i++ {
		nextDiff[i] = diff[i+1] - diff[i]
	}
	return nextDiff
}

func isAllZeroes(diff []int) bool {
	for _, entry := range diff {
		if entry != 0 {
			return false
		}
	}
	return true
}

func extrapolate(history string) int {

	var historyArray []int = make([]int, 0)

	for _, stringNumber := range strings.Split(history, " ") {
		var entry int
		entry, _ = strconv.Atoi(stringNumber)
		historyArray = append(historyArray, entry)
	}

	var diffs [][]int = make([][]int, 0)

	diffs = append(diffs, historyArray)

	diffs = append(diffs, computeNextDiff(historyArray))

	for !isAllZeroes(diffs[len(diffs)-1]) {
		diffs = append(diffs, computeNextDiff(diffs[len(diffs)-1]))
	}

	diffs[len(diffs)-1] = append(diffs[len(diffs)-1], 0)

	for i := len(diffs) - 2; i >= 0; i-- {
		var extrapolatedDiff int = diffs[i][len(diffs[i])-1] + diffs[i+1][len(diffs[i+1])-1]
		diffs[i] = append(diffs[i], extrapolatedDiff)
	}

	return diffs[0][len(diffs[0])-1]
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string = strings.Split(string(input), "\n")

	var sum int = 0

	for _, line := range lines {
		sum += extrapolate(line)
	}

	fmt.Println(sum)
}
