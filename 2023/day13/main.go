package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

func isReflection(input []string, afterIndex int) bool {
	if afterIndex < 0 || afterIndex+1 >= len(input) {
		return false
	}
	for i, j := afterIndex, afterIndex+1; i >= 0 && j < len(input); i, j = i-1, j+1 {
		if debug {
			fmt.Printf("i: %v, j: %v\n", i, j)
		}
		if input[i] != input[j] {
			if debug {
				fmt.Println("Not reflected here")
				fmt.Printf("input[%v]: %v\n", i, input[i])
				fmt.Printf("input[%v]: %v\n", j, input[j])
			}
			return false
		} else if debug {
			fmt.Println("Reflected input")
			fmt.Printf("input[%v]: %v\n", i, input[i])
			fmt.Printf("input[%v]: %v\n", j, input[j])
		}
	}
	if debug {
		fmt.Println("Reflected here\n", afterIndex, afterIndex+1)
	}
	return true
}

func findReflection(pattern []string) int {
	summary := 0
	if debug {
		fmt.Printf("Pattern:\n")
		for _, line := range pattern {
			fmt.Printf("%v\n", line)
		}
	}
	columns := make([]string, len(pattern[0]))
	for x := 0; x < len(pattern[0]); x++ {
		column := ""
		for y := 0; y < len(pattern); y++ {
			column += string(pattern[y][x])
		}
		columns[x] = column
	}
	for x := 0; x < len(columns)-1; x++ {
		if debug {
			fmt.Printf("x: %v–%v\n", x, x+1)
		}
		reflectsHere := isReflection(columns, x)
		if reflectsHere {
			summary += x + 1
			if debug {
				fmt.Printf("Columns reflect here: %v\n", x+1)
			}
		}

	}
	for y := 0; y < len(pattern)-1; y++ {
		if debug {
			fmt.Printf("y: %v–%v\n", y, y+1)
		}
		reflectsHere := isReflection(pattern, y)
		if reflectsHere {
			summary += (y + 1) * 100
			if debug {
				fmt.Printf("Lines reflect here: %v\n", y+1)
			}
		}
	}
	return summary
}

func partOne(lines []string) {
	lines = append(lines, "")
	patternIndex := 0
	patternStartIndex := 0
	sumOfSummaries := 0
	for lineIndex, line := range lines {
		if line == "" {
			summary := findReflection(lines[patternStartIndex:lineIndex])
			if debug {
				fmt.Printf("Summary: %v\n\n", summary)
			}
			sumOfSummaries += summary
			patternIndex++
			patternStartIndex = lineIndex + 1
			continue
		}
	}
	fmt.Printf("Sum of summaries: %v\n", sumOfSummaries)
}

func partTwo(lines []string) {

}

func main() {
	day := 13
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
