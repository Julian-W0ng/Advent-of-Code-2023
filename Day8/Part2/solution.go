package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func isAtEnd(currentNodes *[]string) bool {
	for _, node := range *currentNodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func lcm(a, b int) int {
	var gcd int = 1
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return a * b / gcd
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var directions string = strings.Split(string(input), "\n")[0]

	var lines []string = strings.Split(string(input), "\n")[2:]

	var nodeMap map[string]Node = make(map[string]Node)

	for _, line := range lines {
		var key string = line[:3]
		var left string = line[7:10]
		var right string = line[12:15]
		nodeMap[key] = Node{left, right}
	}

	var currentNodes []string = make([]string, 0)

	for key := range nodeMap {
		if key[2] == 'A' {
			currentNodes = append(currentNodes, key)
		}
	}

	var numberOfIterations []int = make([]int, len(currentNodes))
	var lengthOfDirections int = len(directions)

	for i := range currentNodes {
		for numberOfIterations[i] = 0; currentNodes[i][2] != 'Z'; numberOfIterations[i]++ {
			for j := 0; j < lengthOfDirections; j++ {
				if directions[j] == 'L' {
					currentNodes[i] = nodeMap[currentNodes[i]].left
				} else {
					currentNodes[i] = nodeMap[currentNodes[i]].right
				}
			}
		}
	}

	// Find lcm of all numbers in numberOfIterations
	var numberOfSteps int = numberOfIterations[0]
	for i := 1; i < len(numberOfIterations); i++ {
		numberOfSteps = lcm(numberOfSteps, numberOfIterations[i])
	}
	numberOfSteps *= lengthOfDirections

	fmt.Println(numberOfSteps)
}
