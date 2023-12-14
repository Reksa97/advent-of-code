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
	case "south":
		for y := len(grid) - 2; y >= 0; y-- {
			for x := 0; x < len(grid[y]); x++ {
				if newGrid[y][x] == "O" {
					for yy := y + 1; yy < len(grid) && newGrid[yy][x] == "."; yy++ {
						newGrid[yy][x] = "O"
						newGrid[yy-1][x] = "."
					}
				}
			}
		}
	case "west":
		for x := 1; x < len(grid[0]); x++ {
			for y := 0; y < len(grid); y++ {
				if newGrid[y][x] == "O" {
					for xx := x - 1; xx >= 0 && newGrid[y][xx] == "."; xx-- {
						newGrid[y][xx] = "O"
						newGrid[y][xx+1] = "."
					}
				}
			}
		}
	case "east":
		for x := len(grid[0]) - 2; x >= 0; x-- {
			for y := 0; y < len(grid); y++ {
				if newGrid[y][x] == "O" {
					for xx := x + 1; xx < len(grid[0]) && newGrid[y][xx] == "."; xx++ {
						newGrid[y][xx] = "O"
						newGrid[y][xx-1] = "."
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

type Cycle struct {
	i     int
	cycle int
}

func partTwo(lines []string) {
	cycles := 1000000000
	grid := parseInput(lines)
	differentGrids := make(map[string]Cycle)
	for i := 1; i <= cycles; i++ {
		if i%(cycles/100) == 0 {
			fmt.Println((i*100)/cycles, "%")
		}
		grid = tilt(grid, "north")
		grid = tilt(grid, "west")
		grid = tilt(grid, "south")
		grid = tilt(grid, "east")
		gridString := ""
		for _, line := range grid {
			gridString += fmt.Sprintf("%v", line)
		}
		if cycle, ok := differentGrids[gridString]; ok {
			if debug {
				fmt.Println("Found repeating grid at", i, cycle)
			}
			if i-cycle.i == cycle.cycle {
				if (cycles-i)%cycle.cycle == 0 {
					if debug {
						fmt.Println("Found cycle that end at the end of all cycles", cycle.cycle)
					}
					break
				}
			}
			differentGrids[gridString] = Cycle{i, i - cycle.i}
		} else {
			differentGrids[gridString] = Cycle{i, -1}
		}
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
