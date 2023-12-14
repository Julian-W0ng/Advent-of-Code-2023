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

	var numberOfSteps int
	var length int = len(directions)

	for numberOfSteps = 0; !isAtEnd(&currentNodes); numberOfSteps++ {
		var directionIndex int = numberOfSteps % length
		if directions[directionIndex] == 'L' {
			for i, node := range currentNodes {
				currentNodes[i] = nodeMap[node].left
			}
		} else {
			for i, node := range currentNodes {
				currentNodes[i] = nodeMap[node].right
			}
		}
	}

	fmt.Println(numberOfSteps)
}
