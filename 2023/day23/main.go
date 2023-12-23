package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

type Coordinate struct {
	x int
	y int
}

var paths [][]Coordinate

func canGoTo(position Coordinate, positionToCheck Coordinate, lines []string, pathIndex int) bool {
	if positionToCheck.x < 0 || positionToCheck.y < 0 || positionToCheck.y >= len(lines) || positionToCheck.x >= len(lines[positionToCheck.y]) {
		return false
	}
	nextKey := string(lines[positionToCheck.y][positionToCheck.x])

	if nextKey == "#" {
		return false
	}
	for _, pos := range paths[pathIndex] {
		if pos.x == positionToCheck.x && pos.y == positionToCheck.y {
			return false
		}
	}

	switch nextKey {
	case ".":
		return true
	case ">":
		return position.x < positionToCheck.x
	case "<":
		return position.x > positionToCheck.x
	case "^":
		return position.y > positionToCheck.y
	case "v":
		return position.y < positionToCheck.y
	default:
		panic(fmt.Sprintf("Unknown key: %v", nextKey))
	}
}

func goTo(start Coordinate, end Coordinate, lines []string, pathIndex int) {
	paths[pathIndex] = append(paths[pathIndex], start)

	if start.x == end.x && start.y == end.y {
		return
	}

	branches := 0

	originalPath := make([]Coordinate, len(paths[pathIndex]))
	copy(originalPath, paths[pathIndex])

	// Check if we can go north
	if canGoTo(start, Coordinate{start.x, start.y - 1}, lines, pathIndex) {
		goTo(Coordinate{start.x, start.y - 1}, end, lines, pathIndex)
		branches++
	}

	// Check if we can go south
	if canGoTo(start, Coordinate{start.x, start.y + 1}, lines, pathIndex) {
		nextPathIndex := pathIndex
		if branches > 0 {
			paths = append(paths, make([]Coordinate, len(originalPath)))
			copy(paths[len(paths)-1], originalPath)
			nextPathIndex = len(paths) - 1
		}
		goTo(Coordinate{start.x, start.y + 1}, end, lines, nextPathIndex)
		branches++
	}

	// Check if we can go west
	if canGoTo(start, Coordinate{start.x - 1, start.y}, lines, pathIndex) {
		nextPathIndex := pathIndex
		if branches > 0 {
			paths = append(paths, make([]Coordinate, len(originalPath)))
			copy(paths[len(paths)-1], originalPath)
			nextPathIndex = len(paths) - 1
		}
		goTo(Coordinate{start.x - 1, start.y}, end, lines, nextPathIndex)
		branches++
	}

	// Check if we can go east
	if canGoTo(start, Coordinate{start.x + 1, start.y}, lines, pathIndex) {
		nextPathIndex := pathIndex
		if branches > 0 {
			paths = append(paths, make([]Coordinate, len(originalPath)))
			copy(paths[len(paths)-1], originalPath)
			nextPathIndex = len(paths) - 1
		}
		goTo(Coordinate{start.x + 1, start.y}, end, lines, nextPathIndex)
		branches++
	}

	if debug {
		fmt.Printf("branches: %v, paths length: %v, start: %v, end: %v, pathIndex: %v\n", branches, len(paths), start, end, pathIndex)
	}
}

func printPath(path []Coordinate, lines []string) {
	for _, pos := range path {
		lines[pos.y] = lines[pos.y][:pos.x] + "+" + lines[pos.y][pos.x+1:]
	}
	fmt.Println()
	for _, line := range lines {
		fmt.Printf("%v\n", line)
	}
	fmt.Println()
}

func partOne(lines []string) {
	copyLines := make([]string, len(lines))
	copy(copyLines, lines)
	start := Coordinate{1, 0}
	end := Coordinate{len(lines[0]) - 2, len(lines) - 1}
	paths = make([][]Coordinate, 1)
	paths[0] = make([]Coordinate, 0)
	goTo(start, end, copyLines, 0)
	//fmt.Printf("Longest path: %v\n", longestPath)
	longestPath := 0
	for i, path := range paths {
		pathLength := len(path) - 1 // -1 because we don't count the start
		if debug {
			copy(copyLines, lines)
			fmt.Printf("Path: %v\n", i)
			printPath(path, copyLines)
		}
		if path[pathLength].x != end.x || path[pathLength].y != end.y {
			if debug {
				fmt.Printf("Path didn't reach end: %v\n", path)
			}
			continue
		}
		if pathLength > longestPath {
			if debug {
				fmt.Printf("New longest path: %v, length %v\n", path, pathLength)
			}
			longestPath = pathLength
		} else if debug {
			fmt.Printf("Path too short: %v, length %v\n", path, pathLength)
		}
	}
	fmt.Printf("Longest path: %v\n", longestPath)
}

func partTwo(lines []string) {

}

func main() {
	day := 23
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
