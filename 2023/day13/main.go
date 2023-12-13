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
		if debug && part == 1 {
			fmt.Printf("i: %v, j: %v\n", i, j)
		}
		if input[i] != input[j] {
			if debug && part == 1 {
				fmt.Println("Not reflected here")
				fmt.Printf("input[%v]: %v\n", i, input[i])
				fmt.Printf("input[%v]: %v\n", j, input[j])
			}
			return false
		} else if debug && part == 1 {
			fmt.Println("Reflected input")
			fmt.Printf("input[%v]: %v\n", i, input[i])
			fmt.Printf("input[%v]: %v\n", j, input[j])
		}
	}
	if debug && part == 1 {
		fmt.Println("Reflected here\n", afterIndex, afterIndex+1)
	}
	return true
}

func findReflection(pattern []string, originalSummary int) int {
	summary := 0
	if debug && part == 1 {
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
		if debug && part == 1 {
			fmt.Printf("x: %v–%v\n", x, x+1)
		}
		reflectsHere := isReflection(columns, x)
		if reflectsHere {
			if part == 1 {
				summary += x + 1
			}
			if part == 2 && x+1 != originalSummary {
				summary += x + 1
			}

			if debug && part == 1 {
				fmt.Printf("Columns reflect here: %v\n", x+1)
			}
		}
	}

	for y := 0; y < len(pattern)-1; y++ {
		if debug && part == 1 {
			fmt.Printf("y: %v–%v\n", y, y+1)
		}
		reflectsHere := isReflection(pattern, y)
		if reflectsHere {
			if part == 1 {
				summary += (y + 1) * 100
			}

			if part == 2 && (y+1)*100 != originalSummary {
				summary += (y + 1) * 100
			}
			if debug && part == 1 {
				fmt.Printf("Lines reflect here: %v\n", y+1)
			}
		}
	}
	return summary
}

func findSmudge(pattern []string, originalSummary int) int {
	for x := 0; x < len(pattern[0]); x++ {
		for y := 0; y < len(pattern); y++ {
			newPattern := make([]string, len(pattern))
			copy(newPattern, pattern)
			if pattern[y][x] == '.' {
				newPattern[y] = newPattern[y][:x] + "#" + newPattern[y][x+1:]
			} else {
				newPattern[y] = newPattern[y][:x] + "." + newPattern[y][x+1:]
			}
			summary := findReflection(newPattern, originalSummary)
			if y == 7 && x == len(pattern[0])-1 {
				fmt.Println(pattern[y][x], newPattern[y][x])
				for _, line := range newPattern {
					fmt.Println(line)
				}
				fmt.Println(summary, originalSummary)
			}
			if summary > 0 && summary != originalSummary {
				if debug {
					fmt.Printf("Smudge: %v %v (summary %v)\n", x, y, summary)
				}
				return summary
			}
		}
	}
	for _, line := range pattern {
		fmt.Println(line)
	}
	panic("No smudge found")
}

func findReflectionWithSmudge(pattern []string) int {

	originalSummary := findReflection(pattern, -1)

	return findSmudge(pattern, originalSummary)
}

func solve(lines []string) {
	lines = append(lines, "")
	patternIndex := 0
	patternStartIndex := 0
	sumOfSummaries := 0
	for lineIndex, line := range lines {
		if line == "" {
			summary := 0
			if part == 1 {
				summary = findReflection(lines[patternStartIndex:lineIndex], -1)
			} else {
				summary = findReflectionWithSmudge(lines[patternStartIndex:lineIndex])
			}
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

func main() {
	day := 13
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
