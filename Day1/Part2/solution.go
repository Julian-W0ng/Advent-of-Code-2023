package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// read input.txt file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}

	// split data into lines
	lines := strings.Split(string(data), "\n")

	var sum int = 0
	// array of digits as characters
	charDigits := [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	// array of digits as words
	wordDigits := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range lines {
		firstDigit := -1
		firstIndex := len(line)
		lastDigit := -1
		lastIndex := 0

		// get the first digit of the line
		for i := 0; i < len(line); i++ {
			for j, d := range charDigits {
				if d == line[i] {
					firstDigit = j
					firstIndex = i
					goto foundFirst
				}
			}
		}

	foundFirst:

		for i, d := range wordDigits {
			var index int = strings.Index(line, d)
			if index != -1 && index < firstIndex {
				firstDigit = i
				firstIndex = index
			}
		}

		// get the last digit of the line
		for i := len(line) - 1; i >= 0; i-- {
			for j, d := range charDigits {
				if d == line[i] {
					lastDigit = j
					lastIndex = i
					goto foundLast
				}
			}
		}

	foundLast:

		for i, d := range wordDigits {
			var index int = strings.LastIndex(line, d)
			if index != -1 && index > lastIndex {
				lastDigit = i
				lastIndex = index
			}
		}

		sum += firstDigit*10 + lastDigit
	}

	fmt.Println("Sum of all numbers in the file is: ", sum)
}
