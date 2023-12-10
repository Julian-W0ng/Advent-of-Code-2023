package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	copies := make([]int, len(lines))

	for i := 0; i < len(copies); i++ {
		copies[i] = 1
	}

	for lineIndex, line := range lines {
		myNumbers := strings.Split(strings.Split(line, "|")[1], " ")
		winningNumbers := strings.Split(strings.Split(strings.Split(line, "|")[0], ":")[1], " ")

		// remove empty strings
		for i := 0; i < len(myNumbers); i++ {
			if myNumbers[i] == "" {
				myNumbers = append(myNumbers[:i], myNumbers[i+1:]...)
				i--
			}
		}
		for i := 0; i < len(winningNumbers); i++ {
			if winningNumbers[i] == "" {
				winningNumbers = append(winningNumbers[:i], winningNumbers[i+1:]...)
				i--
			}
		}

		cardPoints := 0

		for _, myNumber := range myNumbers {
			for _, winningNumber := range winningNumbers {
				if myNumber == winningNumber {
					cardPoints++
				}
			}
		}

		for i := lineIndex + 1; i <= lineIndex+cardPoints && i < len(lines); i++ {
			copies[i] += copies[lineIndex]
		}
	}

	sum := 0
	for _, copy := range copies {
		sum += copy
	}

	fmt.Println(sum)
}
