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

const (
	up    = 0
	down  = 1
	left  = 2
	right = 3
)

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

func isEdgePosition(pos Position, matrix [][]byte) bool {
	if pos.y == 0 || pos.y == len(matrix)-1 {
		return true
	}
	if pos.x == 0 || pos.x == len(matrix[pos.y])-1 {
		return true
	}
	return false
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

func addSquare(pos Position, matrix [][]byte, loopSquares [][]bool, squares *[][]bool) {
	if !isPositionValid(pos, matrix) {
		return
	}
	if loopSquares[pos.y][pos.x] {
		return
	}
	if (*squares)[pos.y][pos.x] {
		return
	}

	(*squares)[pos.y][pos.x] = true
	addSquare(Position{pos.x - 1, pos.y}, matrix, loopSquares, squares)
	addSquare(Position{pos.x + 1, pos.y}, matrix, loopSquares, squares)
	addSquare(Position{pos.x, pos.y - 1}, matrix, loopSquares, squares)
	addSquare(Position{pos.x, pos.y + 1}, matrix, loopSquares, squares)
}

func addSquares(pos Position, direction int, matrix [][]byte, loopSquares [][]bool, leftSquares *[][]bool, rightSquares *[][]bool) {
	if direction == up {
		addSquare(Position{pos.x - 1, pos.y}, matrix, loopSquares, leftSquares)
		addSquare(Position{pos.x + 1, pos.y}, matrix, loopSquares, rightSquares)
	} else if direction == down {
		addSquare(Position{pos.x - 1, pos.y}, matrix, loopSquares, rightSquares)
		addSquare(Position{pos.x + 1, pos.y}, matrix, loopSquares, leftSquares)
	} else if direction == left {
		addSquare(Position{pos.x, pos.y + 1}, matrix, loopSquares, leftSquares)
		addSquare(Position{pos.x, pos.y - 1}, matrix, loopSquares, rightSquares)
	} else if direction == right {
		addSquare(Position{pos.x, pos.y + 1}, matrix, loopSquares, rightSquares)
		addSquare(Position{pos.x, pos.y - 1}, matrix, loopSquares, leftSquares)
	}
}

func countSquares(squares [][]bool) int {
	var count int = 0
	for y := 0; y < len(squares); y++ {
		for x := 0; x < len(squares[y]); x++ {
			if squares[y][x] {
				count++
			}
		}
	}
	return count
}

func findAreaInLoop(start Position, matrix [][]byte) int {
	var rightSquares [][]bool = make([][]bool, len(matrix))
	var leftSquares [][]bool = make([][]bool, len(matrix))
	var loopSquares [][]bool = make([][]bool, len(matrix))
	var loopSquares2 [][]bool = make([][]bool, len(matrix))
	for y := 0; y < len(matrix); y++ {
		rightSquares[y] = make([]bool, len(matrix[y]))
		leftSquares[y] = make([]bool, len(matrix[y]))
		loopSquares[y] = make([]bool, len(matrix[y]))
		loopSquares2[y] = make([]bool, len(matrix[y]))
	}

	var currentPos Position = start
	for currentPos != start || !loopSquares[currentPos.y][currentPos.x] {
		loopSquares[currentPos.y][currentPos.x] = true
		if isPositionConnectedUp(currentPos, matrix) && !loopSquares[currentPos.y-1][currentPos.x] {
			currentPos.y--
		} else if isPositionConnectedDown(currentPos, matrix) && !loopSquares[currentPos.y+1][currentPos.x] {
			currentPos.y++
		} else if isPositionConnectedLeft(currentPos, matrix) && !loopSquares[currentPos.y][currentPos.x-1] {
			currentPos.x--
		} else if isPositionConnectedRight(currentPos, matrix) && !loopSquares[currentPos.y][currentPos.x+1] {
			currentPos.x++
		} else {
			currentPos = start
		}
	}

	currentPos = start
	var currentDirection int = up
	if isPositionConnectedUp(currentPos, matrix) {
		currentDirection = up
	} else if isPositionConnectedDown(currentPos, matrix) {
		currentDirection = down
	} else if isPositionConnectedLeft(currentPos, matrix) {
		currentDirection = left
	} else if isPositionConnectedRight(currentPos, matrix) {
		currentDirection = right
	}
	for currentPos != start || !loopSquares2[currentPos.y][currentPos.x] {
		loopSquares2[currentPos.y][currentPos.x] = true
		addSquares(currentPos, currentDirection, matrix, loopSquares, &leftSquares, &rightSquares)

		var nextPos Position = currentPos
		if isPositionConnectedUp(currentPos, matrix) && !loopSquares2[currentPos.y-1][currentPos.x] {
			nextPos.y--
			currentDirection = up
		} else if isPositionConnectedDown(currentPos, matrix) && !loopSquares2[currentPos.y+1][currentPos.x] {
			nextPos.y++
			currentDirection = down
		} else if isPositionConnectedLeft(currentPos, matrix) && !loopSquares2[currentPos.y][currentPos.x-1] {
			nextPos.x--
			currentDirection = left
		} else if isPositionConnectedRight(currentPos, matrix) && !loopSquares2[currentPos.y][currentPos.x+1] {
			nextPos.x++
			currentDirection = right
		} else if isPositionConnectedUp(currentPos, matrix) && currentPos.x == start.x && currentPos.y-1 == start.y {
			nextPos.y--
			currentDirection = up
		} else if isPositionConnectedDown(currentPos, matrix) && currentPos.x == start.x && currentPos.y+1 == start.y {
			nextPos.y++
			currentDirection = down
		} else if isPositionConnectedLeft(currentPos, matrix) && currentPos.x-1 == start.x && currentPos.y == start.y {
			nextPos.x--
			currentDirection = left
		} else if isPositionConnectedRight(currentPos, matrix) && currentPos.x+1 == start.x && currentPos.y == start.y {
			nextPos.x++
			currentDirection = right
		}

		addSquares(currentPos, currentDirection, matrix, loopSquares, &leftSquares, &rightSquares)
		currentPos = nextPos
	}

	var isLeftInside bool = true
	var isRightInside bool = true
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if leftSquares[y][x] && isEdgePosition(Position{x, y}, matrix) {
				isLeftInside = false
			}
			if rightSquares[y][x] && isEdgePosition(Position{x, y}, matrix) {
				isRightInside = false
			}
		}
	}

	if isLeftInside {
		return countSquares(leftSquares)
	} else if isRightInside {
		return countSquares(rightSquares)
	} else {
		return -1
	}
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

	var area int = findAreaInLoop(start, matrix)

	fmt.Println(area)
}
