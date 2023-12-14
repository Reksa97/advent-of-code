package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

func parseInput(lines []string) [][]string {
	grid := make([][]string, len(lines))
	for y, line := range lines {
		grid[y] = make([]string, len(line))
		for x, char := range line {
			grid[y][x] = string(char)
		}
	}
	return grid
}

func tilt(grid [][]string, direction string) [][]string {
	newGrid := make([][]string, len(grid))
	for y := 0; y < len(grid); y++ {
		newGrid[y] = make([]string, len(grid[y]))
		for x := 0; x < len(grid[y]); x++ {
			newGrid[y][x] = grid[y][x]
		}
	}
	switch direction {
	case "north":
		for y := 1; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if newGrid[y][x] == "O" {
					for yy := y - 1; yy >= 0 && newGrid[yy][x] == "."; yy-- {
						newGrid[yy][x] = "O"
						newGrid[yy+1][x] = "."
					}
				}
			}
		}
	default:
		panic("Unknown direction")
	}

	return newGrid
}

func partOne(lines []string) {
	grid := parseInput(lines)
	if debug {
		for _, line := range grid {
			fmt.Println(line)
		}
	}
	grid = tilt(grid, "north")
	if debug {
		fmt.Println("\ntilted north")
	}
	fullLoad := 0
	for y, line := range grid {
		if debug {
			fmt.Println(line)
		}

		for _, char := range line {
			if char == "O" {
				fullLoad += len(grid) - y
			}
		}
	}
	fmt.Printf("Full load: %v\n", fullLoad)
}

func partTwo(lines []string) {

}

func main() {
	day := 14
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
