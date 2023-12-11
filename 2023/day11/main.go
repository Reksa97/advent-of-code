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

func parseInput(lines []string) ([]string, []Galaxy) {
	originalLinesLenght := len(lines)
	addedLines := 0
	for y := 0; y < originalLinesLenght; y++ {
		actualY := y + addedLines
		if debug {
			fmt.Printf("Line %v (%v): %v\n", y, actualY, lines[actualY])
		}
		isEmpty := true
		for _, char := range lines[actualY] {
			if char == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			if debug {
				fmt.Printf("Line %v is empty\n", actualY)
			}
			// duplicate current line
			lines = append(lines, "")
			copy(lines[actualY+1:], lines[actualY:])
			addedLines++
		}
	}

	originalColumnsLenght := len(lines[0])
	addedColumns := 0
	for x := 0; x < originalColumnsLenght; x++ {
		actualX := x + addedColumns

		isEmpty := true
		for y := 0; y < len(lines); y++ {
			if lines[y][actualX] == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			if debug {
				fmt.Printf("Column %v is empty\n", actualX)
				for _, line := range lines {
					fmt.Printf("	%v\n", string(line[actualX]))
				}
			}
			for y := 0; y < len(lines); y++ {
				lines[y] = lines[y][:actualX] + "." + lines[y][actualX:]
			}
			addedColumns++
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
	return lines, galaxies
}

func partOne(lines []string) {
	lines, galaxies := parseInput(lines)
	sumOfDistances := 0
	for i, galaxy := range galaxies {
		if debug {
			fmt.Printf("Galaxy %v: %v\n", i, galaxy)
		}
		for j, otherGalaxy := range galaxies[i+1:] {
			if debug {
				fmt.Printf("	%v: %v\n", j, otherGalaxy)
			}
			absX := galaxy.x - otherGalaxy.x
			if absX < 0 {
				absX = -absX
			}
			absY := galaxy.y - otherGalaxy.y
			if absY < 0 {
				absY = -absY
			}
			shortestPath := absX + absY
			if debug {
				fmt.Printf("		%v\n", shortestPath)
			}
			sumOfDistances += shortestPath
		}
	}
	fmt.Printf("Sum of distances: %v\n", sumOfDistances)
}

func partTwo(lines []string) {

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

	if part == 1 {
		partOne(lines)
	} else {
		partTwo(lines)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
