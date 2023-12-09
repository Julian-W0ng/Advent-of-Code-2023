package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var digits [10]byte = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

type Position struct {
	lineNumber, offset int
}

type Number struct {
	number       int
	isPartNumber bool
}

func findStartOfNumber(p Position, charArray [][]byte, numbers map[Position]Number) Position {
	lineNumber := p.lineNumber
	offset := p.offset
	for offset > 0 && slices.Contains(digits[:], charArray[lineNumber][offset-1]) {
		offset--
	}
	return Position{lineNumber, offset}
}

func isValidPosition(lineNumber int, offset int, numberOfLines int, lineWidth int) bool {
	return lineNumber >= 0 && lineNumber < numberOfLines && offset >= 0 && offset < lineWidth
}

func main() {
	input, error := os.ReadFile("input.txt")
	if error != nil {
		panic(error)
	}

	lines := strings.Split(string(input), "\n")

	numbers := make(map[Position]Number)

	charArray := make([][]byte, len(lines))
	for i := range charArray {
		charArray[i] = make([]byte, len(lines[i]))
	}

	for lineIndex, line := range lines {
		for offsetIndex, char := range line {
			charArray[lineIndex][offsetIndex] = byte(char)
		}
	}

	for lineIndex, line := range charArray {
		beginningOffset := -1
		lastWasDigit := false
		for offsetIndex, char := range line {
			if slices.Contains(digits[:], char) {
				for i, digit := range digits {
					if digit == char {
						if lastWasDigit {
							numbers[Position{lineIndex, beginningOffset}] =
								Number{numbers[Position{lineIndex, beginningOffset}].number*10 + i, false}
						} else {
							beginningOffset = offsetIndex
							numbers[Position{lineIndex, beginningOffset}] = Number{i, false}
						}
					}
				}
				lastWasDigit = true
			} else {
				lastWasDigit = false
			}
		}
	}

	radius := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	sum := 0

	for lineIndex, line := range charArray {
		for offsetIndex, char := range line {
			if char != '.' && !slices.Contains(digits[:], char) {
				partsFound := make(map[Position]int)
				for _, r := range radius {
					if isValidPosition(lineIndex+r[0], offsetIndex+r[1], len(charArray), len(charArray[0])) {
						if slices.Contains(digits[:], charArray[lineIndex+r[0]][offsetIndex+r[1]]) {
							startOfNumber := findStartOfNumber(Position{lineIndex + r[0], offsetIndex + r[1]}, charArray, numbers)
							numbers[startOfNumber] = Number{numbers[startOfNumber].number, true}
							partsFound[startOfNumber] = numbers[startOfNumber].number
						}
					}
				}
				if len(partsFound) == 2 && char == '*' {
					product := 1
					for _, value := range partsFound {
						product *= value
					}
					sum += product
				}
			}
		}
	}

	fmt.Println(sum)
}
