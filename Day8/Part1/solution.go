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
		if byte(node[2]) != 'Z' {
			return false
		}
	}
	return true
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

	var currentNode string = "AAA"

	var numberOfSteps int

	for numberOfSteps = 0; currentNode != "ZZZ"; numberOfSteps++ {
		var directionIndex int = numberOfSteps % len(directions)
		if directions[directionIndex] == 'L' {
			currentNode = nodeMap[currentNode].left
		} else {
			currentNode = nodeMap[currentNode].right
		}
	}

	fmt.Println(numberOfSteps)
}
