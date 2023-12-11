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
			comingFromUp := y > 0 && (lines[y-1][x] == '|' || lines[y-1][x] == 'F' || lines[y-1][x] == '7')
			comingFromDown := y < len(lines)-1 && (lines[y+1][x] == '|' || lines[y+1][x] == 'J' || lines[y+1][x] == 'L')
			comingFromLeft := x > 0 && (lines[y][x-1] == '-' || lines[y][x-1] == 'L' || lines[y][x-1] == 'F')
			comingFromRight := x < len(lines[0])-1 && (lines[y][x+1] == '-' || lines[y][x+1] == 'J' || lines[y][x+1] == '7')
			if char == 'S' {
				start = Position{x: x, y: y}
				if comingFromUp && comingFromDown {
					lines[y] = lines[y][:x] + "|" + lines[y][x+1:]
				} else if comingFromLeft && comingFromRight {
					lines[y] = lines[y][:x] + "-" + lines[y][x+1:]
				} else if comingFromUp && comingFromLeft {
					lines[y] = lines[y][:x] + "J" + lines[y][x+1:]
				} else if comingFromUp && comingFromRight {
					lines[y] = lines[y][:x] + "L" + lines[y][x+1:]
				} else if comingFromDown && comingFromLeft {
					lines[y] = lines[y][:x] + "7" + lines[y][x+1:]
				} else if comingFromDown && comingFromRight {
					lines[y] = lines[y][:x] + "F" + lines[y][x+1:]
				} else {
					panic(fmt.Sprintf("Invalid start position: %v", start))
				}
			}
		}
	}
	return start, distances
}

func getExpandedInput(lines []string) []string {
	expandedLines := []string{}
	for y, line := range lines {
		expandedLines = append(expandedLines, "", "", "")
		for _, char := range line {
			switch char {
			case '.':
				expandedLines[y*3] += "..."
				expandedLines[y*3+1] += "..."
				expandedLines[y*3+2] += "..."
			case '|':
				expandedLines[y*3] += ".|."
				expandedLines[y*3+1] += ".|."
				expandedLines[y*3+2] += ".|."
			case '-':
				expandedLines[y*3] += "..."
				expandedLines[y*3+1] += "---"
				expandedLines[y*3+2] += "..."
			case 'L':
				expandedLines[y*3] += ".|."
				expandedLines[y*3+1] += ".L-"
				expandedLines[y*3+2] += "..."
			case 'J':
				expandedLines[y*3] += ".|."
				expandedLines[y*3+1] += "-J."
				expandedLines[y*3+2] += "..."
			case 'F':
				expandedLines[y*3] += "..."
				expandedLines[y*3+1] += ".F-"
				expandedLines[y*3+2] += ".|."
			case '7':
				expandedLines[y*3] += "..."
				expandedLines[y*3+1] += "-7."
				expandedLines[y*3+2] += ".|."
			case 'S':
				expandedLines[y*3] += "..."
				expandedLines[y*3+1] += ".S."
				expandedLines[y*3+2] += "..."
			}
		}
	}
	return expandedLines
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

func expandOuterDots(lines []string, position Position) {
	if position.x < 0 || position.x >= len(lines[0]) || position.y < 0 || position.y >= len(lines) {
		return
	}
	char := lines[position.y][position.x]
	if char == '.' {
		lines[position.y] = lines[position.y][:position.x] + "0" + lines[position.y][position.x+1:]
		expandOuterDots(lines, Position{x: position.x - 1, y: position.y})     // left
		expandOuterDots(lines, Position{x: position.x + 1, y: position.y})     // right
		expandOuterDots(lines, Position{x: position.x, y: position.y - 1})     // up
		expandOuterDots(lines, Position{x: position.x, y: position.y + 1})     // down
		expandOuterDots(lines, Position{x: position.x - 1, y: position.y - 1}) // left and up
		expandOuterDots(lines, Position{x: position.x - 1, y: position.y + 1}) // left and down
		expandOuterDots(lines, Position{x: position.x + 1, y: position.y - 1}) // right and up
		expandOuterDots(lines, Position{x: position.x + 1, y: position.y + 1}) // right and down
	}
}

func partTwo(lines []string) {
	startPosition, distances := parseInput(lines)
	traverse(lines, startPosition, distances, -1)
	for y, line := range distances {
		for x, distance := range line {
			if distance == -1 {
				lines[y] = lines[y][:x] + "." + lines[y][x+1:]
			}
		}
	}
	lines = getExpandedInput(lines)
	if debug {
		for _, line := range lines {
			fmt.Println(line)
		}
	}
	for x, char := range lines[0] {
		if char == '.' {
			expandOuterDots(lines, Position{x: x, y: 0})
		}
	}
	for x, char := range lines[len(lines)-1] {
		if char == '.' {
			expandOuterDots(lines, Position{x: x, y: len(lines) - 1})
		}
	}
	for y, line := range lines {
		if line[0] == '.' {
			expandOuterDots(lines, Position{x: 0, y: y})
		}
		if line[len(line)-1] == '.' {
			expandOuterDots(lines, Position{x: len(line) - 1, y: y})
		}
	}
	dotsInsideLoop := 0
	for y := 1; y < len(lines)-1; y += 3 {
		for x := 1; x < len(lines[0])-1; x += 3 {
			if lines[y-1][x-1] == '.' && lines[y-1][x] == '.' && lines[y-1][x+1] == '.' &&
				lines[y][x-1] == '.' && lines[y][x] == '.' && lines[y][x+1] == '.' &&
				lines[y+1][x-1] == '.' && lines[y+1][x] == '.' && lines[y+1][x+1] == '.' {
				dotsInsideLoop++
			}
		}
	}
	fmt.Printf("Answer: %v\n", dotsInsideLoop)
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
