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

	points := 0

	for _, line := range lines {
		cardPoints := 0
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

		for _, myNumber := range myNumbers {
			for _, winningNumber := range winningNumbers {
				if myNumber == winningNumber {
					if cardPoints == 0 {
						cardPoints = 1
					} else {
						cardPoints *= 2
					}
				}
			}
		}
		points += cardPoints
	}

	fmt.Println(points)
}
