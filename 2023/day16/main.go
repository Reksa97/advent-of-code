package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

func goTo(x int, y int, direction string, lines []string, visited [][]map[string]bool) {
	if x < 0 || y < 0 || y >= len(lines) || x >= len(lines[y]) {
		return
	}
	if visited[y][x][direction] {
		return
	}
	visited[y][x][direction] = true
	currentKey := string(lines[y][x])
	if debug {
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Printf("x: %v, y: %v, direction: %v, currentKey: %v\n", x, y, direction, currentKey)
	}

	switch currentKey {
	case ".", ">", "<", "^", "v":
		lines[y] = lines[y][:x] + direction + lines[y][x+1:]
		if direction == ">" {
			goTo(x+1, y, ">", lines, visited)
		} else if direction == "<" {
			goTo(x-1, y, "<", lines, visited)
		} else if direction == "^" {
			goTo(x, y-1, "^", lines, visited)
		} else if direction == "v" {
			goTo(x, y+1, "v", lines, visited)
		}
	case "|":
		if direction == ">" || direction == "<" {
			goTo(x, y-1, "^", lines, visited)
			goTo(x, y+1, "v", lines, visited)
		} else if direction == "^" {
			goTo(x, y-1, "^", lines, visited)
		} else if direction == "v" {
			goTo(x, y+1, "v", lines, visited)
		}
	case "-":
		if direction == "^" || direction == "v" {
			goTo(x-1, y, "<", lines, visited)
			goTo(x+1, y, ">", lines, visited)
		} else if direction == "<" {
			goTo(x-1, y, "<", lines, visited)
		} else if direction == ">" {
			goTo(x+1, y, ">", lines, visited)
		}

	case "/":
		if direction == "^" {
			goTo(x+1, y, ">", lines, visited)
		} else if direction == "v" {
			goTo(x-1, y, "<", lines, visited)
		} else if direction == "<" {
			goTo(x, y+1, "v", lines, visited)
		} else if direction == ">" {
			goTo(x, y-1, "^", lines, visited)
		}
	case "\\":
		if direction == "^" {
			goTo(x-1, y, "<", lines, visited)
		} else if direction == "v" {
			goTo(x+1, y, ">", lines, visited)
		} else if direction == "<" {
			goTo(x, y-1, "^", lines, visited)
		} else if direction == ">" {
			goTo(x, y+1, "v", lines, visited)
		}
	}
}

func partOne(lines []string, firstX int, firstY int, firstDirection string) int {
	visited := make([][]map[string]bool, len(lines))
	for i, line := range lines {
		if debug {
			fmt.Printf("i: %v, line: %v\n", i, line)
		}

		arrayOfMaps := make([]map[string]bool, len(line))
		for i := range arrayOfMaps {
			arrayOfMaps[i] = make(map[string]bool)
		}

		visited[i] = append(visited[i], arrayOfMaps...)
	}

	goTo(firstX, firstY, firstDirection, lines, visited)

	energized := 0

	for _, line := range visited {
		for _, cell := range line {
			if len(cell) > 0 {
				if debug {
					fmt.Printf("#")
				}

				energized++
			} else if debug {
				fmt.Printf(".")
			}
		}
		if debug {
			fmt.Printf("\n")
		}
	}
	if debug {
		fmt.Printf("\n")
		for _, line := range lines {
			fmt.Println(line)
		}
	}

	if part == 1 {
		fmt.Printf("Energized: %v\n", energized)
	}
	return energized
}

func partTwo(lines []string) {
	originalLines := make([]string, len(lines))
	copy(originalLines, lines)
	maximumEnergized := 0
	for x := 0; x < len(lines[0]); x++ {
		lines = make([]string, len(originalLines))
		copy(lines, originalLines)
		energized := partOne(lines, x, 0, "v")
		if energized > maximumEnergized {
			maximumEnergized = energized
		}
		lines = make([]string, len(originalLines))
		copy(lines, originalLines)
		energized = partOne(lines, x, len(lines)-1, "^")
		if energized > maximumEnergized {
			maximumEnergized = energized
		}
	}
	for y := 0; y < len(lines); y++ {
		lines = make([]string, len(originalLines))
		copy(lines, originalLines)
		energized := partOne(lines, 0, y, ">")
		if energized > maximumEnergized {
			maximumEnergized = energized
		}
		lines = make([]string, len(originalLines))
		copy(lines, originalLines)
		energized = partOne(lines, len(lines[0])-1, y, "<")
		if energized > maximumEnergized {
			maximumEnergized = energized
		}
	}
	fmt.Printf("Maximum energized: %v\n", maximumEnergized)
}

func main() {
	day := 16
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Part: %v\n", part)

	startTime := time.Now()

	if part == 1 {
		partOne(lines, 0, 0, ">")
	} else {
		partTwo(lines)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
