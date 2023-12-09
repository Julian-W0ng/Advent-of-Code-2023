package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const RED int = 0
const GREEN int = 1
const BLUE int = 2

type Game struct {
	rounds [][3]int
}

func main() {
	var numCubes [3]int
	numCubes[RED] = 12
	numCubes[GREEN] = 13
	numCubes[BLUE] = 14
	var stringColours [3]string
	stringColours[RED] = "red"
	stringColours[GREEN] = "green"
	stringColours[BLUE] = "blue"

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	games := make([]Game, len(lines))

	// Parse input
	for _, line := range lines {
		gameID, error := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		if error != nil {
			fmt.Println("Error converting gameID to int")
			panic(error)
		}
		gameID--
		gameRounds := strings.Split(strings.Split(line, ":")[1], ";")
		for _, round := range gameRounds {
			roundCount := [3]int{0, 0, 0}
			cubes := strings.Split(round, ",")
			for _, cube := range cubes {
				for j, colour := range stringColours {
					if strings.Contains(cube, colour) {
						count, error := strconv.Atoi(strings.Split(cube, " ")[1])
						if error != nil {
							fmt.Println("Error converting cube to int")
							panic(error)
						}
						roundCount[j] = count
					}
				}
			}
			games[gameID].rounds = append(games[gameID].rounds, roundCount)
		}
	}

	// Check how many games are valid
	var powerSum int = 0

	for _, game := range games {
		minRed, minGreen, minBlue := 0, 0, 0
		for _, round := range game.rounds {
			if round[RED] > minRed {
				minRed = round[RED]
			}
			if round[GREEN] > minGreen {
				minGreen = round[GREEN]
			}
			if round[BLUE] > minBlue {
				minBlue = round[BLUE]
			}
		}
		powerSum += minRed * minGreen * minBlue
	}

	fmt.Println(powerSum)
}
