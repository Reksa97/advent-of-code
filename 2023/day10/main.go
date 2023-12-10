package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

type Position struct {
	x int
	y int
}

func parseInput(lines []string) (Position, [][]int) {
	distances := [][]int{}
	start := Position{}
	for y, line := range lines {
		distances = append(distances, []int{})

		for x, char := range line {
			distances[y] = append(distances[y], -1)
			if char == 'S' {
				start = Position{x: x, y: y}
			}
		}
	}
	return start, distances
}

func traverse(lines []string, position Position, distances [][]int, prevDistance int) {
	if debug {
		fmt.Printf("traverse position: %v, prevDistance: %v\n", position, prevDistance)
	}
	if position.x < 0 || position.x >= len(lines[0]) || position.y < 0 || position.y >= len(lines) {
		if debug {
			fmt.Printf("Out of bounds: %v\n", position)
		}
		return
	}
	distance := distances[position.y][position.x]
	hasBeenVisited := distance != -1 && distance <= prevDistance
	if hasBeenVisited {
		if debug {
			fmt.Printf("Already visited %v\n", position)
		}
		return
	}
	distance = prevDistance + 1
	distances[position.y][position.x] = distance
	char := lines[position.y][position.x]

	switch char {
	case '|':
		traverse(lines, Position{x: position.x, y: position.y - 1}, distances, distance) // up
		traverse(lines, Position{x: position.x, y: position.y + 1}, distances, distance) // down
	case '-':
		traverse(lines, Position{x: position.x - 1, y: position.y}, distances, distance) // left
		traverse(lines, Position{x: position.x + 1, y: position.y}, distances, distance) // right
	case 'L':
		traverse(lines, Position{x: position.x, y: position.y - 1}, distances, distance) // up
		traverse(lines, Position{x: position.x + 1, y: position.y}, distances, distance) // right
	case 'J':
		traverse(lines, Position{x: position.x, y: position.y - 1}, distances, distance) // up
		traverse(lines, Position{x: position.x - 1, y: position.y}, distances, distance) // left
	case 'F':
		traverse(lines, Position{x: position.x, y: position.y + 1}, distances, distance) // down
		traverse(lines, Position{x: position.x + 1, y: position.y}, distances, distance) // right
	case '7':
		traverse(lines, Position{x: position.x - 1, y: position.y}, distances, distance) // left
		traverse(lines, Position{x: position.x, y: position.y + 1}, distances, distance) // down
	case 'S':
		if position.y > 0 {
			upChar := lines[position.y-1][position.x]
			if upChar == '|' || upChar == 'F' || upChar == '7' {
				traverse(lines, Position{x: position.x, y: position.y - 1}, distances, distance) // up
			}
		}

		if position.y < len(lines)-1 {
			downChar := lines[position.y+1][position.x]
			if downChar == '|' || downChar == 'J' || downChar == 'L' {
				traverse(lines, Position{x: position.x, y: position.y + 1}, distances, distance) // down
			}
		}

		if position.x > 0 {
			leftChar := lines[position.y][position.x-1]
			if leftChar == '-' || leftChar == 'L' || leftChar == 'F' {
				traverse(lines, Position{x: position.x - 1, y: position.y}, distances, distance) // left
			}
		}

		if position.x < len(lines[0])-1 {
			rightChar := lines[position.y][position.x+1]
			if rightChar == '-' || rightChar == 'J' || rightChar == '7' {
				traverse(lines, Position{x: position.x + 1, y: position.y}, distances, distance) // right
			}
		}
	}
}

func partOne(lines []string) {
	startPosition, distances := parseInput(lines)
	fmt.Printf("startPosition: %v\n", startPosition)
	if debug {
		fmt.Printf("distances: %v\n", distances)
	}
	traverse(lines, startPosition, distances, -1)
	max := 0
	for _, line := range distances {
		if debug {
			fmt.Println(line)
		}
		for _, distance := range line {
			if distance > max {
				max = distance
			}
		}
	}
	fmt.Printf("Answer: %v\n", max)
}

func partTwo(lines []string) {

}

func main() {
	day := 10
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Part: %v\n", part)

	startTime := time.Now()

	if part == 1 {
		partOne(lines)
	} else {
		partTwo(lines)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
