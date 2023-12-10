package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const STARTOFDESTINATION int = 0
const STARTOFSOURCE int = 1
const RANGE int = 2

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	seeds := strings.Split(strings.Split(strings.Split(string(input), "\n")[0], ": ")[1], " ")

	sections := strings.Split(string(input), "\n\n")[1:]

	sectionsArray := make([][][3]int, len(sections))

	for i, section := range sections {
		lines := strings.Split(section, "\n")[1:]

		sectionsArray[i] = make([][3]int, len(lines))

		for j, line := range lines {
			numbers := strings.Split(line, " ")

			sectionsArray[i][j][STARTOFDESTINATION], _ = strconv.Atoi(numbers[STARTOFDESTINATION])
			sectionsArray[i][j][STARTOFSOURCE], _ = strconv.Atoi(numbers[STARTOFSOURCE])
			sectionsArray[i][j][RANGE], _ = strconv.Atoi(numbers[RANGE])
		}
	}

	min := math.MaxInt64
	for _, seed := range seeds {
		translation, _ := strconv.Atoi(seed)

		for _, section := range sectionsArray {
			for _, line := range section {
				if translation >= line[STARTOFSOURCE] && translation < line[STARTOFSOURCE]+line[RANGE] {
					translation = line[STARTOFDESTINATION] + (translation - line[STARTOFSOURCE])
					break
				}
			}
		}

		if translation < min {
			min = translation
		}
	}

	fmt.Println(min)
}
