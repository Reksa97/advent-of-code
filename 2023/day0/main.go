package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"time"
)

func partOne(lines []string, debug bool) {

}

func partTwo(lines []string, debug bool) {

}

func main() {
	day := 0
	lines, err := common.ReadInput(day, os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	part := common.GetPart(os.Args)
	debug := common.IsTestMode(os.Args)
	fmt.Printf("Part: %v\n", part)

	startTime := time.Now()

	if part == 1 {
		partOne(lines, debug)
	} else {
		partTwo(lines, debug)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v μs (%v ms)\n", duration.Microseconds(), duration.Milliseconds())
}
