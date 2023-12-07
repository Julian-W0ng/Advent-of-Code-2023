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
	// constant array of digits as words
	// digitsWords := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range lines {
		firstDigit := -1
		// firstDigitIndex := len(line)
		lastDigit := -1
		// lastDigitIndex := 0

		found := false

		// get the first digit of the line
		for i := 0; i < len(line) && !found; i++ {
			for j, d := range charDigits {
				if d == line[i] {
					firstDigit = j
					// firstDigitIndex = i
					found = true
					break
				}
			}
		}

		found = false

		// get the last digit of the line
		for i := len(line) - 1; i >= 0 && !found; i-- {
			for j, d := range charDigits {
				if d == line[i] {
					lastDigit = j
					// firstDigitIndex = i
					found = true
					break
				}
			}
		}

		sum += firstDigit*10 + lastDigit

	}

	fmt.Println("Sum of all numbers in the file is: ", sum)
}
