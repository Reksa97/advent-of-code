package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

type Galaxy struct {
	x, y int
}

func parseInput(lines []string) ([]Galaxy, []bool, []bool) {
	expandedLines := make([]bool, len(lines))
	expandedColumns := make([]bool, len(lines[0]))
	for y := 0; y < len(lines); y++ {
		if debug {
			fmt.Printf("Line %v: %v\n", y, lines[y])
		}
		isEmpty := true
		for _, char := range lines[y] {
			if char == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			if debug {
				fmt.Printf("Line %v is empty\n", y)
			}
			expandedLines[y] = true
		}
	}

	for x := 0; x < len(lines[0]); x++ {

		isEmpty := true
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			if debug {
				fmt.Printf("Column %v is empty\n", x)
				for _, line := range lines {
					fmt.Printf("	%v\n", string(line[x]))
				}
			}
			expandedColumns[x] = true
		}
	}

	galaxies := []Galaxy{}
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Galaxy{x, y})
			}
		}
	}

	if debug {
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println(galaxies)
	}
	return galaxies, expandedColumns, expandedLines
}

func solve(lines []string) {
	EXPAND_MULTIPLIER := 2
	if part == 2 {
		EXPAND_MULTIPLIER = 1_000_000
	}
	galaxies, expandedColumns, expandedLines := parseInput(lines)
	sumOfDistances := 0
	for i, galaxy := range galaxies {
		if debug {
			fmt.Printf("Galaxy %v: %v\n", i, galaxy)
		}
		for j, otherGalaxy := range galaxies[i+1:] {
			if debug {
				fmt.Printf("	%v: %v\n", j, otherGalaxy)
			}
			xDirection := 1
			if otherGalaxy.x < galaxy.x {
				xDirection = -1
			}
			yDirection := 1
			if otherGalaxy.y < galaxy.y {
				yDirection = -1
			}
			shortestPath := 0
			if galaxy.x != otherGalaxy.x {
				for x := galaxy.x + xDirection; x != otherGalaxy.x; x += xDirection {
					if debug {
						fmt.Printf("		%v\n", x)
					}
					if !expandedColumns[x] {
						shortestPath++
					} else {
						shortestPath += EXPAND_MULTIPLIER
					}
				}
				shortestPath++
			}
			if galaxy.y != otherGalaxy.y {
				for y := galaxy.y + yDirection; y != otherGalaxy.y; y += yDirection {
					if !expandedLines[y] {
						shortestPath++
					} else {
						shortestPath += EXPAND_MULTIPLIER
					}
				}
				shortestPath++
			}

			if debug {
				fmt.Printf("		%v\n", shortestPath)
			}
			sumOfDistances += shortestPath
		}
	}
	fmt.Printf("Sum of distances: %v\n", sumOfDistances)
}

func main() {
	day := 11
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Part: %v\n", part)

	startTime := time.Now()

	solve(lines)

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
