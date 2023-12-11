package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

const TIME int = 0
const DISTANCE int = 1

var wg sync.WaitGroup

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// trim double spaces
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' && input[i+1] == ' ' {
			input = append(input[:i], input[i+1:]...)
			i--
		}
	}

	lines := strings.Split(string(input), "\n")

	times := strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ")

	distances := strings.Split(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " ")

	timeDistances := make([][2]int, len(times))

	for i := 0; i < len(times); i++ {
		timeDistances[i][TIME], _ = strconv.Atoi(times[i])
		timeDistances[i][DISTANCE], _ = strconv.Atoi(distances[i])
	}

	var waysToWin = make(chan int, len(timeDistances))

	for i := 0; i < len(timeDistances); i++ {
		wg.Add(1)
		go findNumberOfWaysToWin(timeDistances[i][TIME], timeDistances[i][DISTANCE], waysToWin)
	}

	wg.Wait()

	product := 1
	for i := 0; i < len(timeDistances); i++ {
		product *= <-waysToWin
	}

	println(product)
}

func findNumberOfWaysToWin(time int, distance int, waysToWin chan int) {
	defer wg.Done()

	numberOfWaysToWin := 0

	for i := 0; i <= time; i++ {
		var distanceTravelled int = i * (time - i)
		if distanceTravelled > distance {
			numberOfWaysToWin++
		}
	}

	waysToWin <- numberOfWaysToWin
}
