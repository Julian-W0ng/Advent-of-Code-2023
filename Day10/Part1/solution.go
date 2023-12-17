package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

var connectedUp []byte = []byte{'S', '|', 'L', 'J'}
var connectedDown []byte = []byte{'S', '|', '7', 'F'}
var connectedLeft []byte = []byte{'S', '-', 'J', '7'}
var connectedRight []byte = []byte{'S', '-', 'F', 'L'}

func isPositionValid(pos Position, matrix [][]byte) bool {
	if pos.y < 0 || pos.y >= len(matrix) {
		return false
	}
	if pos.x < 0 || pos.x >= len(matrix[pos.y]) {
		return false
	}
	return true
}

func isPositionConnectedUp(pos Position, matrix [][]byte) bool {
	var up Position = Position{pos.x, pos.y - 1}
	if !isPositionValid(up, matrix) || !isPositionValid(pos, matrix) {
		return false
	}
	for _, upChar := range connectedUp {
		if matrix[pos.y][pos.x] == upChar {
			for _, downChar := range connectedDown {
				if matrix[up.y][up.x] == downChar {
					return true
				}
			}
		}
	}
	return false
}

func isPositionConnectedDown(pos Position, matrix [][]byte) bool {
	var down Position = Position{pos.x, pos.y + 1}
	if !isPositionValid(down, matrix) || !isPositionValid(pos, matrix) {
		return false
	}
	for _, downChar := range connectedDown {
		if matrix[pos.y][pos.x] == downChar {
			for _, upChar := range connectedUp {
				if matrix[down.y][down.x] == upChar {
					return true
				}
			}
		}
	}
	return false
}

func isPositionConnectedLeft(pos Position, matrix [][]byte) bool {
	var left Position = Position{pos.x - 1, pos.y}
	if !isPositionValid(left, matrix) || !isPositionValid(pos, matrix) {
		return false
	}
	for _, leftChar := range connectedLeft {
		if matrix[pos.y][pos.x] == leftChar {
			for _, rightChar := range connectedRight {
				if matrix[left.y][left.x] == rightChar {
					return true
				}
			}
		}
	}
	return false
}

func isPositionConnectedRight(pos Position, matrix [][]byte) bool {
	var right Position = Position{pos.x + 1, pos.y}
	if !isPositionValid(right, matrix) || !isPositionValid(pos, matrix) {
		return false
	}
	for _, rightChar := range connectedRight {
		if matrix[pos.y][pos.x] == rightChar {
			for _, leftChar := range connectedLeft {
				if matrix[right.y][right.x] == leftChar {
					return true
				}
			}
		}
	}
	return false
}

func findMaxDepth(start Position, matrix [][]byte) int {

	var queue []Position = []Position{start}
	var visited [][]bool = make([][]bool, len(matrix))

	for y := range visited {
		visited[y] = make([]bool, len(matrix[y]))
	}

	var depth int = -1

	for len(queue) > 0 {
		fmt.Println(queue)
		var nextQueue []Position
		for _, pos := range queue {
			if visited[pos.y][pos.x] {
				continue
			}
			visited[pos.y][pos.x] = true
			if isPositionConnectedUp(pos, matrix) && !visited[pos.y-1][pos.x] {
				nextQueue = append(nextQueue, Position{pos.x, pos.y - 1})
			}
			if isPositionConnectedDown(pos, matrix) && !visited[pos.y+1][pos.x] {
				nextQueue = append(nextQueue, Position{pos.x, pos.y + 1})
			}
			if isPositionConnectedLeft(pos, matrix) && !visited[pos.y][pos.x-1] {
				nextQueue = append(nextQueue, Position{pos.x - 1, pos.y})
			}
			if isPositionConnectedRight(pos, matrix) && !visited[pos.y][pos.x+1] {
				nextQueue = append(nextQueue, Position{pos.x + 1, pos.y})
			}
		}
		queue = nextQueue
		depth++
	}

	for y := range visited {
		fmt.Println(visited[y])
	}

	return depth
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string = strings.Split(string(input), "\n")

	var start Position

	var matrix [][]byte = make([][]byte, len(lines))

	for y, line := range lines {
		var row []byte = make([]byte, len(line))
		for x, char := range line {
			row[x] = byte(char)
			if char == 'S' {
				start.x = x
				start.y = y
			}
		}
		matrix[y] = row
	}

	var maxDepth int = findMaxDepth(start, matrix)

	fmt.Println(maxDepth)
}
