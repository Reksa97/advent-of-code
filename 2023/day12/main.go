package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"strings"
	"time"
)

var part = common.GetPart(os.Args)
var debug = common.IsTestMode(os.Args)

func countPossibleArrangements(conditions string, pattern []int, patternIndex int, adjacentDamagedCount int, previousConditions string) int {

	/* if debug {
		fmt.Printf("\nConditions: %v|%v\n", previousConditions, conditions)
		fmt.Printf("Pattern index: %v\n", patternIndex)
		fmt.Printf("Adjacent damaged count: %v\n", adjacentDamagedCount)
	} */

	if len(conditions) == 0 {
		if (patternIndex == len(pattern)-1 && adjacentDamagedCount == pattern[patternIndex]) || (patternIndex == len(pattern) && adjacentDamagedCount == 0) {
			return 1
		}
		return 0
	}

	switch conditions[0] {
	case '.':
		// operational
		if adjacentDamagedCount > 0 && patternIndex < len(pattern) && pattern[patternIndex] != adjacentDamagedCount {
			// incorrect amount of damaged next to each other
			return 0
		}
		newPatternIndex := patternIndex
		if (patternIndex < len(pattern)) && (pattern[patternIndex] == adjacentDamagedCount && adjacentDamagedCount > 0) {
			newPatternIndex++
		}
		return countPossibleArrangements(conditions[1:], pattern, newPatternIndex, 0, previousConditions+".")
	case '#':
		// damaged
		if patternIndex >= len(pattern) || pattern[patternIndex] == adjacentDamagedCount {
			// no more damaged sequences allowed or too many damaged next to each other
			return 0
		}
		return countPossibleArrangements(conditions[1:], pattern, patternIndex, adjacentDamagedCount+1, previousConditions+"#")
	case '?':
		// unknown
		return countPossibleArrangements("."+conditions[1:], pattern, patternIndex, adjacentDamagedCount, previousConditions) +
			countPossibleArrangements("#"+conditions[1:], pattern, patternIndex, adjacentDamagedCount, previousConditions)
	}
	panic(fmt.Sprintf("Unknown condition: %v", conditions[0]))
}

func partOne(lines []string) {
	totalArrangements := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		conditions := fields[0]
		pattern := common.ConvertToInt(strings.Split(fields[1], ","))
		if debug {
			fmt.Printf("New line Conditions: %v\n", conditions)
			fmt.Printf("New line Pattern: %v\n", pattern)
		}
		possibleArrangmenets := countPossibleArrangements(conditions, pattern, 0, 0, "")
		if debug {
			fmt.Printf("\nLine calculated Possible arrangements: %v\n\n", possibleArrangmenets)
		}
		totalArrangements += possibleArrangmenets
	}
	fmt.Printf("Total arrangements: %v\n", totalArrangements)
}

func partTwo(lines []string) {

}

func main() {
	day := 12
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
