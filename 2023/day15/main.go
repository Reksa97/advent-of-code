package main

import (
	"aoc2023/common"
	"fmt"
	"os"
	"strconv"
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

		if debug && part == 1 {
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

type Lens struct {
	label string
	value int
}

func partTwo(lines []string) {
	statements := strings.Split(lines[0], ",")
	// 256 boxes
	boxes := make([][]Lens, 256)
	for _, statement := range statements {
		statementSplit := strings.Split(statement, "=")
		label := ""
		if len(statementSplit) == 1 {
			operation := statement[len(statement)-1]
			label = statement[:len(statement)-1]
			boxIndex := hash(label)
			for i, box := range boxes[boxIndex] {
				if box.label == label {
					// remove item from boxes[boxIndex] if lens with the label is in the box
					// and the operation is -
					if operation == '-' {
						boxes[boxIndex] = append(boxes[boxIndex][:i], boxes[boxIndex][i+1:]...)
					}
					break
				}
			}

			if debug {
				fmt.Printf("\nbox:%v, op:%v label:%v\n", boxIndex, string(operation), label)
			}
			continue
		}
		label = statementSplit[0]
		boxIndex := hash(label)
		value, _ := strconv.Atoi(statementSplit[1])
		foundAndReplaced := false
		newLens := Lens{label: label, value: value}
		for i, box := range boxes[boxIndex] {
			if box.label == label {
				// replace item in boxes[boxIndex] if lens with the label is in the box
				firstHalf := append(boxes[boxIndex][:i], newLens)
				boxes[boxIndex] = append(firstHalf, boxes[boxIndex][i+1:]...)
				foundAndReplaced = true
				break
			}
		}
		if !foundAndReplaced {
			boxes[boxIndex] = append(boxes[boxIndex], newLens)
		}
		if debug {
			fmt.Printf("\nbox:%v, %v = %v\n", boxIndex, label, value)
			for i, box := range boxes {
				if len(box) > 0 {
					fmt.Println(i, box)
				}
			}
		}
	}
	focusingPower := 0
	for i, box := range boxes {
		if len(box) > 0 {
			if debug {
				fmt.Println(i, box)
			}
			for j, lens := range box {
				lensPower := (1 + i) * (1 + j) * lens.value
				focusingPower += lensPower
			}
		}
	}
	fmt.Println(focusingPower)
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
