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

func hash(str string) int {
	currentValue := 0
	for _, char := range str {
		currentValue += int(char)
		multiplied := currentValue * 17
		divided := multiplied / 256
		currentValue = multiplied - divided*256

		if debug {
			fmt.Printf("\n%v (%v - %v*256) %v", char, multiplied, divided, currentValue)
		}
	}
	return currentValue
}

func partOne(lines []string) {
	sum := 0
	statements := strings.Split(lines[0], ",")
	for _, statement := range statements {
		hashed := hash(statement)
		sum += hashed
		if debug {
			fmt.Printf("\n%v = %v\n", statement, hashed)
		}
	}
	fmt.Println(sum)
}

func partTwo(lines []string) {

}

func main() {
	day := 15
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
