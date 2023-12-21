package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

func parseInput(lines []string) ([]string, int, int) {
	newLines := make([]string, 0)
	for _, line := range lines {
		newLines = append(newLines, line)
	}
	for y, line := range newLines {
		for x, char := range line {
			if char == 'S' {
				newLines = append(newLines[:y], append([]string{line[:x] + "." + line[x+1:]}, newLines[y+1:]...)...)
				return newLines, x, y
			}
		}
	}
	panic("Start not found")
}

type Coordinate struct {
	x int
	y int
}

func partOne(originalLines []string) {
	steps := 64
	lines, startX, startY := parseInput(originalLines)
	if debug {
		fmt.Println(startX, startY)
		for _, line := range lines {
			fmt.Println(line)
		}
		for _, line := range originalLines {
			fmt.Println(line)
		}
	}

	possiblePlaces := make(map[Coordinate]bool)
	possiblePlaces[Coordinate{startX, startY}] = true
	for step := 1; step <= steps; step++ {
		nextPlaces := make(map[Coordinate]bool, 0)
		for coordinate := range possiblePlaces {
			// South
			if coordinate.y+1 < len(lines) && lines[coordinate.y+1][coordinate.x] == '.' {
				nextPlaces[Coordinate{coordinate.x, coordinate.y + 1}] = true
			}
			// North
			if coordinate.y-1 >= 0 && lines[coordinate.y-1][coordinate.x] == '.' {
				nextPlaces[Coordinate{coordinate.x, coordinate.y - 1}] = true
			}
			// East
			if coordinate.x+1 < len(lines[coordinate.y]) && lines[coordinate.y][coordinate.x+1] == '.' {
				nextPlaces[Coordinate{coordinate.x + 1, coordinate.y}] = true
			}
			// West
			if coordinate.x-1 >= 0 && lines[coordinate.y][coordinate.x-1] == '.' {
				nextPlaces[Coordinate{coordinate.x - 1, coordinate.y}] = true
			}
		}
		if debug {
			fmt.Println("\nStep", step)
			for y, line := range lines {
				for x, char := range line {
					if _, ok := nextPlaces[Coordinate{x, y}]; ok {
						fmt.Print("O")
					} else {
						fmt.Print(string(char))
					}
				}
				fmt.Println()
			}
		}
		possiblePlaces = nextPlaces
	}

	fmt.Println("Possible plots:", len(possiblePlaces))

}

func partTwo(lines []string) {

}

func main() {
	day := 21
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
